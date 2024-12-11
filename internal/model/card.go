package model

type Card struct {
	Name        string         // 卡牌名称
	Type        string         // 卡牌类型（例如：攻击、治疗等）
	Energy      int            // 使用该卡牌所需的能量
	Target      string         // 目标（例如：enemy, self）
	Block       int            // 卡牌的护甲值
	Damage      int            // 卡牌的伤害值
	Power       map[string]int // 卡牌的攻击力或特殊能力
	Action      []string       // 卡牌的动作（例如：apply poison, heal）
	Description string         // 卡牌的描述
	Image       string         // 卡牌的图像或相关图片
	Upgrade     int            // 卡牌是否升级
}

func (Card) TableName() string {
	return "card"
}
