package model

// WSPlayer 表示玩家状态
type WSPlayer struct {
	NewCard []string       `json:"newCard"`
	HP      int            `json:"hp"`
	Block   int            `json:"block"`
	Energy  int            `json:"energy"`
	Power   map[string]int `json:"power"` // 支持多个状态效果
	Image   string         `json:"image"`
}

// WSMonster 表示怪物状态
type WSMonster struct {
	HP          int            `json:"hp"`
	Block       int            `json:"block"`
	Power       map[string]int `json:"power"`       // 支持多个状态效果
	ActionName  string         `json:"actionName"`  // 支持多种行为效果
	ActionValue int            `json:"actionValue"` // 行为效果的值
	Image       string         `json:"image"`
}

// Message 表示 WebSocket 消息的总结构
type Message struct {
	Type      string     `json:"type"`
	Name      string     `json:"name"` // 仅当 Type="card" 时使用
	WSPlayer  *WSPlayer  `json:"player"`
	WSMonster *WSMonster `json:"monster"`
}

func newMessage(t string) *Message {
	return &Message{
		Type: t,
	}
}
