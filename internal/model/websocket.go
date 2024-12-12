package model

// WSPlayer 表示玩家状态
type WSPlayer struct {
	NewCard []string       `json:"newCard,omitempty"`
	HP      int            `json:"hp,omitempty"`
	Block   int            `json:"block,omitempty"`
	Energy  int            `json:"energy,omitempty"`
	Power   map[string]int `json:"power,omitempty"` // 支持多个状态效果
}

// WSMonster 表示怪物状态
type WSMonster struct {
	HP          int            `json:"hp,omitempty"`
	Block       int            `json:"block,omitempty"`
	Power       map[string]int `json:"power,omitempty"`       // 支持多个状态效果
	ActionName  string         `json:"actionName,omitempty"`  // 支持多种行为效果
	ActionValue int            `json:"actionValue,omitempty"` // 行为效果的值
}

// Message 表示 WebSocket 消息的总结构
type Message struct {
	Type      string     `json:"type"`
	Name      string     `json:"name,omitempty"` // 仅当 Type="card" 时使用
	WSPlayer  *WSPlayer  `json:"player,omitempty"`
	WSMonster *WSMonster `json:"monster,omitempty"`
}

func newMessage(t string) *Message {
	return &Message{
		Type: t,
	}
}
