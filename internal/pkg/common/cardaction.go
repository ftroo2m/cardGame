package common

import "cardGame/internal/model"

func DealDamageEqualToBlock(player *model.Player, monster *model.Monster) {
	monster.HP -= DamageCalculateToMonster(player, monster, player.Block)
}

func DrawCards(game *model.Game, num int) {
	game.DrawCard(num)
}

func AddEnergyToPlayer(player *model.Player, energy int) {
	player.Energy += energy
}

func RemovePlayerDebuffs(player *model.Player) {
	player.Power["weak"] = 0
	player.Power["vulnerable"] = 0
	player.Power["poison"] = 0
}

func RemoveHealth(player *model.Player, damage int) {
	player.HP -= damage
}

func AddRegenToDamage(player *model.Player, monster *model.Monster, damage int) {
	player.Power["regen"] += damage
	monster.HP -= DamageCalculateToMonster(player, monster, damage)
}
