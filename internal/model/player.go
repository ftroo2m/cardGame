package model

// Player 结构体表示一个玩家
type Player struct {
	ID     string         // 玩家的唯一标识符
	HP     int            // 生命值
	Block  int            // 护甲值，用于防止伤害
	Energy int            // 能量，玩家每回合可用的能量
	Power  map[string]int // 攻击力或特殊能力值
	Cards  []Card         // 玩家的手牌
	Image  string         // 玩家头像或卡牌图像
}

func NewPlayer(id string, hp, block, energy int, power map[string]int, cards []Card, image string) *Player {
	return &Player{
		ID:     id,
		HP:     hp,
		Block:  block,
		Energy: energy,
		Power:  power,
		Cards:  cards,
		Image:  image,
	}
}
