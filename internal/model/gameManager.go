package model

import (
	"cardGame/config"
	"cardGame/ent"
	"cardGame/ent/userconfig"
	"context"
	"encoding/json"
	"github.com/gorilla/websocket"
	"sync"
)

type GameManager struct {
	games map[string]*Game
	mu    sync.Mutex
}

var GameManagerUse = NewGameManager()

func NewGameManager() *GameManager {
	return &GameManager{
		games: make(map[string]*Game),
	}
}

func (gm *GameManager) CreateGame(playerID string, conn *websocket.Conn) {
	gm.mu.Lock()
	defer gm.mu.Unlock()
	var player *Player
	userconfig, _ := config.SqlClient.UserConfig.Query().Where(userconfig.PlayerID(playerID)).First(context.Background())
	player = NewPlayer(playerID, userconfig.PlayerHP, 0, 3, map[string]int{}, []string{}, "")
	var nodes []Node
	json.Unmarshal([]byte(userconfig.Ladder), &nodes)
	var monster ent.Monster
	for i := 0; i < len(nodes); i++ {
		if nodes[i].Visit == 0 {
			monster = config.MonsterList[nodes[i].Name]
			monster.Power = map[string]int{}
			break
		}
	}
	game := NewGame(player, &monster, userconfig.Cards, conn)
	gm.games[playerID] = game
	go game.Run()
}

func (gm *GameManager) GetGame(gameID string) (*Game, bool) {
	gm.mu.Lock()
	defer gm.mu.Unlock()
	game, exists := gm.games[gameID]
	return game, exists
}

func (gm *GameManager) RemoveGame(gameID string) {
	gm.mu.Lock()
	defer gm.mu.Unlock()
	game, exists := gm.games[gameID]
	if exists {
		_ = game.PlayerConn.Close()
	}
	delete(gm.games, gameID)
}
