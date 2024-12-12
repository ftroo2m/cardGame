package model

import (
	"cardGame/config"
	"cardGame/ent"
	"cardGame/ent/userconfig"
	"context"
	"github.com/gorilla/websocket"
	"sync"
)

type GameManager struct {
	games map[string]*Game
	mu    sync.Mutex
}

func NewGameManager() *GameManager {
	return &GameManager{
		games: make(map[string]*Game),
	}
}

func (gm *GameManager) CreateGame(playerID string, monsterName string, conn *websocket.Conn) {
	gm.mu.Lock()
	defer gm.mu.Unlock()
	var player *Player
	userconfig, _ := config.SqlClient.UserConfig.Query().Where(userconfig.PlayerID(playerID)).First(context.Background())
	if userconfig == nil {
		player = NewPlayer(playerID, 50, 0, 3, map[string]int{}, []ent.Card{}, "")
	} else {
		player = NewPlayer(playerID, userconfig.PlayerHP, 0, 3, map[string]int{}, []ent.Card{}, "")
	}
	monster := config.MonsterList[monsterName]
	game := NewGame(player, &monster, conn)
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
