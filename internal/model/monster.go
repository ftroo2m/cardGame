package model

type Monster struct {
	Name    string
	HP      int            // 生命值
	Block   int            // 护甲值
	Power   map[string]int // 攻击力或特殊能力
	Actions map[string]int // 怪物的动作列表
	Image   string         // 怪物的图像
}

func (Monster) TableName() string {
	return "monster"
}
