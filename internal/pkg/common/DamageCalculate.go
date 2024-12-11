package common

import (
	"cardGame/ent"
	"cardGame/internal/model"
)

func DamageCalculateToMonster(player *model.Player, monster *ent.Monster, origin int) int {
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

func DamageCalculateToPlayer(player model.Player, monster ent.Monster, origin int) int {
	// 计算怪物对玩家造成的伤害
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
