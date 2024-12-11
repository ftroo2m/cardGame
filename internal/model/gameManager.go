package model

import (
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

func (gm *GameManager) CreateGame(playerID string, conn *websocket.Conn) *Game {
	gm.mu.Lock()
	defer gm.mu.Unlock()
	player := NewPlayer(playerID, 50, 0, 3, map[string]int{}, []Card{}, "")
	game := NewGame(player, conn)
	gm.games[playerID] = game
	return game
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
