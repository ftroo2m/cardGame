package model

import (
	"cardGame/ent"
)

func DealDamageEqualToBlock(player *Player, monster *ent.Monster) {
	monster.HP -= DamageCalculateToMonster(player, monster, player.Block)
}

func DrawCards(game *Game, num int) {
	game.DrawCard(num)
}

func AddEnergyToPlayer(player *Player, energy int) {
	player.Energy += energy
}

func RemovePlayerDebuffs(player *Player) {
	player.Power["weak"] = 0
	player.Power["vulnerable"] = 0
	player.Power["poison"] = 0
}

func RemoveHealth(player *Player, damage int) {
	player.HP -= damage
}

func AddRegenToDamage(player *Player, monster *ent.Monster, damage int) {
	player.Power["regen"] += damage
	monster.HP -= DamageCalculateToMonster(player, monster, damage)
}

func PowerCause(player *Player, monster *ent.Monster, power string, num int, goal int) {
	if goal == 1 {
		switch power {
		case "regen":
			player.Power["regen"] += num
		case "strength":
			player.Power["strength"] += num
		case "weak":
			monster.Power["weak"] += num
		case "vulnerable":
			monster.Power["vulnerable"] += num
		case "poison":
			monster.Power["poison"] += num
		}
	} else {
		switch power {
		case "regen":
			monster.Power["regen"] += num
		case "strength":
			monster.Power["strength"] += num
		case "weak":
			player.Power["weak"] += num
		case "vulnerable":
			player.Power["vulnerable"] += num
		case "poison":
			player.Power["poison"] += num
		}
	}
}

func PowerCalculatePlayer(player *Player) {
	if player.Power["regen"] != 0 {
		player.HP += player.Power["regen"]
		player.Power["regen"] -= 1
	}
	if player.Power["poison"] != 0 {
		player.HP -= player.Power["poison"]
		player.Power["poison"] -= 1
	}
	if player.Power["strength"] != 0 {
		player.Power["strength"] -= 1
	}
	if player.Power["weak"] != 0 {
		player.Power["weak"] -= 1
	}
	if player.Power["vulnerable"] != 0 {
		player.Power["vulnerable"] -= 1
	}
}

func PowerCalculateMonster(monster *ent.Monster) {
	if monster.Power["regen"] != 0 {
		monster.HP += monster.Power["regen"]
		monster.Power["regen"] -= 1
	}
	if monster.Power["poison"] != 0 {
		monster.HP -= monster.Power["poison"]
		monster.Power["poison"] -= 1
	}
	if monster.Power["strength"] != 0 {
		monster.Power["strength"] -= 1
	}
	if monster.Power["weak"] != 0 {
		monster.Power["weak"] -= 1
	}
	if monster.Power["vulnerable"] != 0 {
		monster.Power["vulnerable"] -= 1
	}
}
