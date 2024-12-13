package model

import (
	"cardGame/config"
	"cardGame/ent"
	"cardGame/ent/monster"
	"cardGame/ent/userconfig"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"sort"
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
	usercf, _ := config.SqlClient.UserConfig.Query().Where(userconfig.PlayerID(g.Player.ID)).First(context.Background())
	var nodes []Node
	json.Unmarshal([]byte(usercf.Ladder), &nodes)
	for i := 0; i < len(nodes); i++ {
		if nodes[i].Name == g.Monster.Name {
			nodes[i].Visit = 1
			break
		}
	}
	jsonWay, _ := json.Marshal(nodes)
	keys := make([]string, 0, len(config.CardList))
	for key := range config.CardList {
		keys = append(keys, key)
	}
	randomKey := keys[g.rng.Intn(len(keys))]
	randomCard := config.CardList[randomKey].Name
	usercf.Cards = append(usercf.Cards, randomCard)
	config.SqlClient.UserConfig.UpdateOneID(usercf.ID).SetLadder(string(jsonWay)).SetCards(usercf.Cards).Save(context.Background())
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
	_, _ = config.SqlClient.UserConfig.Delete().Where(userconfig.PlayerID(g.Player.ID)).Exec(context.Background())
	way := GetWay()
	initialHand := []string{"袭击", "袭击", "袭击", "袭击", "袭击", "防御", "防御", "防御", "防御", "猛击"}
	config.SqlClient.UserConfig.Create().SetPlayerID(g.Player.ID).SetPlayerHP(72).SetPlayerEnergy(3).SetLadder(way).SetCards(initialHand).Save(context.Background())
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

func DealDamageEqualToBlock(player *Player, monster *ent.Monster) {
	monster.HP -= DamageCalculateToMonster(player, monster, player.Block)
}

func DrawCards(game *Game, num int) {
	game.DrawCard(num)
}

func AddEnergyToPlayer(player *Player, energy int) {
	player.Energy += energy
}

func RemovePlayerDebuffs(player *Player) {
	player.Power["weak"] = 0
	player.Power["vulnerable"] = 0
	player.Power["poison"] = 0
}

func RemoveHealth(player *Player, damage int) {
	player.HP -= damage
}

func AddRegenToDamage(player *Player, monster *ent.Monster, damage int) {
	player.Power["regen"] += damage
	monster.HP -= DamageCalculateToMonster(player, monster, damage)
}

func PowerCause(player *Player, monster *ent.Monster, power string, num int, goal int) {
	if goal == 1 {
		switch power {
		case "regen":
			player.Power["regen"] += num
		case "strength":
			player.Power["strength"] += num
		case "weak":
			monster.Power["weak"] += num
		case "vulnerable":
			monster.Power["vulnerable"] += num
		case "poison":
			monster.Power["poison"] += num
		}
	} else {
		switch power {
		case "regen":
			monster.Power["regen"] += num
		case "strength":
			monster.Power["strength"] += num
		case "weak":
			player.Power["weak"] += num
		case "vulnerable":
			player.Power["vulnerable"] += num
		case "poison":
			player.Power["poison"] += num
		}
	}
}

func PowerCalculatePlayer(player *Player) {
	if player.Power["regen"] != 0 {
		player.HP += player.Power["regen"]
		player.Power["regen"] -= 1
	}
	if player.Power["poison"] != 0 {
		player.HP -= player.Power["poison"]
		player.Power["poison"] -= 1
	}
	if player.Power["strength"] != 0 {
		player.Power["strength"] -= 1
	}
	if player.Power["weak"] != 0 {
		player.Power["weak"] -= 1
	}
	if player.Power["vulnerable"] != 0 {
		player.Power["vulnerable"] -= 1
	}
}

func PowerCalculateMonster(monster *ent.Monster) {
	if monster.Power["regen"] != 0 {
		monster.HP += monster.Power["regen"]
		monster.Power["regen"] -= 1
	}
	if monster.Power["poison"] != 0 {
		monster.HP -= monster.Power["poison"]
		monster.Power["poison"] -= 1
	}
	if monster.Power["strength"] != 0 {
		monster.Power["strength"] -= 1
	}
	if monster.Power["weak"] != 0 {
		monster.Power["weak"] -= 1
	}
	if monster.Power["vulnerable"] != 0 {
		monster.Power["vulnerable"] -= 1
	}
}

func DamageCalculateToMonster(player *Player, monster *ent.Monster, origin int) int {
	if origin == 0 {
		return 0
	}
	// 计算玩家对怪物造成的伤害
	damage := origin + player.Power["strength"]
	if monster.Power["vulnerable"] != 0 {
		damage = int(float64(damage) * 1.5)
	}

	if player.Power["weak"] != 0 {
		damage = int(float64(damage) * 0.75)
	}

	if monster.Block > damage {
		monster.Block -= damage
		return 0
	} else {
		damage -= monster.Block
		monster.Block = 0
		return damage
	}
}

func DamageCalculateToPlayer(player *Player, monster *ent.Monster, origin int) int {
	// 计算怪物对玩家造成的伤害

	if origin == 0 {
		return 0
	}

	damage := origin + monster.Power["strength"]

	if player.Power["vulnerable"] != 0 {
		damage = int(float64(damage) * 1.5)
	}
	if monster.Power["weak"] != 0 {
		damage = int(float64(damage) * 0.75)
	}

	if player.Block > damage {
		player.Block -= damage
		return 0
	} else {
		damage -= player.Block
		player.Block = 0
		return damage
	}
}

func GetWay() string {
	way := make([]Node, 8)
	rand.Seed(time.Now().UnixNano())
	id := [7]int{}
	p := 0
	for i := 1; i <= 3; i++ {
		sid := rand.Intn(5) + 1
		for j := 1; j < i; j++ {
			if sid == id[j] {
				p = 1
				break
			}
		}
		if p == 1 {
			i--
			p = 0
		} else {
			id[i] = sid
		}
	}
	for i := 4; i <= 5; i++ {
		sid := rand.Intn(3) + 7
		for j := 4; j < i; j++ {
			if sid == id[j] {
				p = 1
				break
			}
		}
		if p == 1 {
			i--
			p = 0
		} else {
			id[i] = sid
		}
	}
	id[6] = rand.Intn(2) + 10
	sort.Ints(id[:])
	for i := 1; i <= 3; i++ {
		println(id[i])
		monster, err := config.SqlClient.Monster.Query().Where(monster.ID(id[i])).First(context.Background())
		if err != nil {
			fmt.Println("Monster query error:", err)
			println("monster query error")
		}
		println(monster.Name)
		way[i-1] = Node{
			ID:    i - 1,
			Name:  monster.Name,
			Visit: 0,
		}
	}
	way[3] = Node{
		ID:    3,
		Name:  "campfire",
		Visit: 0,
	}
	for i := 4; i <= 5; i++ {
		monster, _ := config.SqlClient.Monster.Query().Where(monster.ID(id[i])).First(context.Background())
		way[i] = Node{
			ID:    i,
			Name:  monster.Name,
			Visit: 0,
		}
	}
	way[6] = Node{
		ID:    6,
		Name:  "campfire",
		Visit: 0,
	}
	monster, _ := config.SqlClient.Monster.Query().Where(monster.ID(id[6])).First(context.Background())
	way[7] = Node{
		ID:    7,
		Name:  monster.Name,
		Visit: 0,
	}
	jsonWay, _ := json.Marshal(way)
	return string(jsonWay)
}
