package model

import (
	"github.com/gorilla/websocket"
	"math/rand"
	"time"
)

type Game struct {
	Player     *Player
	Monster    Monster
	Cards      []Card
	UsedCards  []Card
	PlayerConn *websocket.Conn
	rng        *rand.Rand // 随机数生成器
}

func (g *Game) DrawCard(num int) {
	if num > len(g.Cards)+len(g.UsedCards) {
		num = len(g.Cards) + len(g.UsedCards)
	}
	for i := 0; i < num; i++ {
		if len(g.Cards) == 0 {
			g.Cards = g.UsedCards
			g.UsedCards = []Card{}
		}
		index := g.rng.Intn(len(g.Cards))
		selectedCard := g.Cards[index]
		g.Cards = append(g.Cards[:index], g.Cards[index+1:]...)
		g.Player.Cards = append(g.Player.Cards, selectedCard)
	}
}

func NewGame(player *Player, conn *websocket.Conn) *Game {
	return &Game{
		Player:     player,
		Cards:      []Card{},
		UsedCards:  []Card{},
		PlayerConn: conn,
		rng:        rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}
