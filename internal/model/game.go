package model

import (
	"cardGame/config"
	"cardGame/ent"
	"cardGame/internal/pkg/common"
	"encoding/json"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

type Game struct {
	Player             *Player
	Monster            *ent.Monster
	Cards              []ent.Card
	UsedCards          []ent.Card
	PlayerConn         *websocket.Conn
	rng                *rand.Rand // 随机数生成器
	monsterActionIndex int
	NewCards           []string
}

func (g *Game) DrawCard(num int) {
	if num > len(g.Cards)+len(g.UsedCards) {
		num = len(g.Cards) + len(g.UsedCards)
	}
	for i := 0; i < num; i++ {
		if len(g.Cards) == 0 {
			g.Cards = g.UsedCards
			g.UsedCards = []ent.Card{}
		}
		index := g.rng.Intn(len(g.Cards))
		selectedCard := g.Cards[index]
		g.NewCards = append(g.NewCards, selectedCard.Name)
		g.Cards = append(g.Cards[:index], g.Cards[index+1:]...)
		g.Player.Cards = append(g.Player.Cards, selectedCard)
	}
}

func NewGame(player *Player, monster *ent.Monster, conn *websocket.Conn) *Game {
	return &Game{
		Player:     player,
		Monster:    monster,
		Cards:      []ent.Card{},
		UsedCards:  []ent.Card{},
		PlayerConn: conn,
		rng:        rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (g *Game) Run() {
	g.monsterActionIndex = 0
	for {
		_, message, err := g.PlayerConn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
		}
		var msg Message
		err = json.Unmarshal(message, &msg)
		if err != nil {
			log.Printf("error decoding JSON: %v", err)
			continue
		}
		switch msg.Type {
		case "card":
			g.useCard(msg)
		case "init":
			g.init()
		case "endRound":
			g.endRound()
		}
	}
}

func (g *Game) exit() {
	config.GameManager.RemoveGame(g.Player.ID)
}

func (g *Game) useCard(msg Message) {
	cardName := msg.Name
	card := config.CardList[cardName]
	g.Player.Energy -= card.Energy
	g.Player.Block += card.Block
	common.DamageCalculateToMonster(g.Player, g.Monster, card.Damage)
	for _, action := range card.Action {
		switch action {
		case "removePlayerDebuffs":
			common.RemovePlayerDebuffs(g.Player)
		case "addRegenEqualToAllDamage":
			common.AddRegenToDamage(g.Player, g.Monster, 2)
		case "removeHealth":
			common.RemoveHealth(g.Player, 2)
		case "dealDamageEqualToBlock":
			common.DealDamageEqualToBlock(g.Player, g.Monster)
		case "drawCards":
			common.DrawCards(g, 1)
		case "addEnergy":
			common.AddEnergyToPlayer(g.Player, 1)
		}
	}

	for power, num := range card.Power {
		common.PowerCause(g.Player, g.Monster, power, num, 1)
	}
	g.UsedCards = append(g.UsedCards, card)
	index := 0
	for i, c := range g.Player.Cards {
		if c.Name == cardName {
			index = i
			break
		}
	}
	g.Player.Cards = append(g.Player.Cards[:index], g.Player.Cards[index+1:]...)
	if g.Monster.HP <= 0 {
		g.monsterDeath()
	} else if g.Player.HP <= 0 {
		g.playerDeath()
	} else {
		g.returnMessage("card")
	}
}

func (g *Game) monsterDeath() {
	m := newMessage("monsterDeath")
	js, _ := json.Marshal(m)
	g.PlayerConn.WriteMessage(websocket.TextMessage, js)
	g.exit()
}

func (g *Game) init() {
	g.DrawCard(5)
	g.returnMessage("initReturn")
}

func (g *Game) endRound() {
	common.PowerCalculateMonster(g.Monster)
	action := g.Monster.ActionName[g.monsterActionIndex]
	switch action {
	case "damage":
		common.DamageCalculateToPlayer(g.Player, g.Monster, g.Monster.ActionValue[g.monsterActionIndex])
	case "block":
		g.Monster.Block += g.Monster.ActionValue[g.monsterActionIndex]
	case "weak":
		g.Player.Power["weak"] += g.Monster.ActionValue[g.monsterActionIndex]
	case "vulnerable":
		g.Player.Power["vulnerable"] += g.Monster.ActionValue[g.monsterActionIndex]
	}
	common.PowerCalculatePlayer(g.Player)

	if g.Monster.HP <= 0 {
		g.monsterDeath()
	} else if g.Player.HP <= 0 {
		g.playerDeath()
	} else {
		g.monsterActionIndex = (g.monsterActionIndex + 1) % len(g.Monster.ActionName)
		g.returnMessage("endRound")
	}
}

func (g *Game) playerDeath() {
	m := newMessage("playerDeath")
	js, _ := json.Marshal(m)
	g.PlayerConn.WriteMessage(websocket.TextMessage, js)
	g.exit()
}

func (g *Game) returnMessage(code string) {
	m := newMessage(code)
	m.WSPlayer = &WSPlayer{
		NewCard: g.NewCards,
		HP:      g.Player.HP,
		Block:   g.Player.Block,
		Energy:  g.Player.Energy,
		Power:   g.Player.Power,
	}
	m.WSMonster = &WSMonster{
		HP:          g.Monster.HP,
		Block:       g.Monster.Block,
		Power:       g.Monster.Power,
		ActionName:  g.Monster.ActionName[g.monsterActionIndex],
		ActionValue: g.Monster.ActionValue[g.monsterActionIndex],
	}
	m.Name = code
	js, _ := json.Marshal(m)
	g.PlayerConn.WriteMessage(websocket.TextMessage, js)
}
