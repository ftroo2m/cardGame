package handles

import (
	"cardGame/config"
	"cardGame/ent/leaderboard"
	"cardGame/ent/userconfig"
	"cardGame/internal/util"
	"context"
	"github.com/gin-gonic/gin"
)

func Win(c *gin.Context) {
	cookieName := "playerID"
	cookieValue, _ := c.Cookie(cookieName)
	player, _ := config.SqlClient.Leaderboard.Query().Where(leaderboard.PlayerID(cookieValue)).First(context.Background())
	if player == nil {
		config.SqlClient.Leaderboard.Create().SetPlayerID(cookieValue).SetCounts(1).Save(context.Background())
	} else {
		config.SqlClient.Leaderboard.Update().Where(leaderboard.PlayerID(cookieValue)).SetCounts(player.Counts + 1).Save(context.Background())
	}
	way := util.GetWay()
	initialHand := []string{"袭击", "袭击", "袭击", "袭击", "袭击", "防御", "防御", "防御", "防御", "猛击"}
	config.SqlClient.UserConfig.Delete().Where(userconfig.PlayerID(cookieValue)).Exec(context.Background())
	config.SqlClient.UserConfig.Create().SetPlayerID(cookieValue).SetPlayerHP(72).SetPlayerEnergy(3).SetLadder(way).SetCards(initialHand).SetPlayerID(cookieValue).Save(context.Background())
	util.SuccessResp(c, gin.H{"message": "win"})
}
