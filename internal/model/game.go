package model

import (
	"cardGame/config"
	"cardGame/ent"
	"encoding/json"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

type Game struct {
	Player             *Player
	Monster            *ent.Monster
	Cards              []string
	UsedCards          []string
	PlayerConn         *websocket.Conn
	rng                *rand.Rand // 随机数生成器
	MonsterActionIndex int
	NewCards           []string
}

func (g *Game) DrawCard(num int) {
	if num > len(g.Cards)+len(g.UsedCards) {
		num = len(g.Cards) + len(g.UsedCards)
	}
	for i := 0; i < num; i++ {
		if len(g.Cards) == 0 {
			g.Cards = g.UsedCards
			g.UsedCards = []string{}
		}
		index := g.rng.Intn(len(g.Cards))
		selectedCard := g.Cards[index]
		g.NewCards = append(g.NewCards, selectedCard)
		g.Cards = append(g.Cards[:index], g.Cards[index+1:]...)
		g.Player.Cards = append(g.Player.Cards, selectedCard)
	}
}

func NewGame(player *Player, monster *ent.Monster, conn *websocket.Conn) *Game {
	return &Game{
		Player:     player,
		Monster:    monster,
		Cards:      []string{},
		UsedCards:  []string{},
		PlayerConn: conn,
		rng:        rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (g *Game) Run() {
	g.MonsterActionIndex = 0
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
	GameManagerUse.RemoveGame(g.Player.ID)
}

func (g *Game) useCard(msg Message) {
	cardName := msg.Name
	card := config.CardList[cardName]
	g.Player.Energy -= card.Energy
	g.Player.Block += card.Block
	DamageCalculateToMonster(g.Player, g.Monster, card.Damage)
	for _, action := range card.Action {
		switch action {
		case "removePlayerDebuffs":
			RemovePlayerDebuffs(g.Player)
		case "addRegenEqualToAllDamage":
			AddRegenToDamage(g.Player, g.Monster, 2)
		case "removeHealth":
			RemoveHealth(g.Player, 2)
		case "dealDamageEqualToBlock":
			DealDamageEqualToBlock(g.Player, g.Monster)
		case "drawCards":
			DrawCards(g, 1)
		case "addEnergy":
			AddEnergyToPlayer(g.Player, 1)
		}
	}

	for power, num := range card.Power {
		PowerCause(g.Player, g.Monster, power, num, 1)
	}
	g.UsedCards = append(g.UsedCards, card.Name)
	index := 0
	for i, name := range g.Player.Cards {
		if name == cardName {
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
	PowerCalculateMonster(g.Monster)
	action := g.Monster.ActionName[g.MonsterActionIndex]
	switch action {
	case "damage":
		DamageCalculateToPlayer(g.Player, g.Monster, g.Monster.ActionValue[g.MonsterActionIndex])
	case "block":
		g.Monster.Block += g.Monster.ActionValue[g.MonsterActionIndex]
	case "weak":
		g.Player.Power["weak"] += g.Monster.ActionValue[g.MonsterActionIndex]
	case "vulnerable":
		g.Player.Power["vulnerable"] += g.Monster.ActionValue[g.MonsterActionIndex]
	}
	PowerCalculatePlayer(g.Player)
	g.DrawCard(1)
	if g.Monster.HP <= 0 {
		g.monsterDeath()
	} else if g.Player.HP <= 0 {
		g.playerDeath()
	} else {
		g.MonsterActionIndex = (g.MonsterActionIndex + 1) % len(g.Monster.ActionName)
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
		ActionName:  g.Monster.ActionName[g.MonsterActionIndex],
		ActionValue: g.Monster.ActionValue[g.MonsterActionIndex],
	}
	m.Name = code
	js, _ := json.Marshal(m)
	g.PlayerConn.WriteMessage(websocket.TextMessage, js)
}
