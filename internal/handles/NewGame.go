package handles

import (
	"cardGame/config"
	"cardGame/ent/userconfig"
	"cardGame/internal/model"
	"cardGame/internal/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

func NewGame(c *gin.Context) {
	cookieName := "playerID"
	cookieValue, _ := c.Cookie(cookieName)
	way := model.GetWay()
	initialHand := []string{"袭击", "袭击", "袭击", "袭击", "袭击", "防御", "防御", "防御", "防御", "猛击"}
	config.SqlClient.UserConfig.Delete().Where(userconfig.PlayerID(cookieValue)).Exec(context.Background())
	config.SqlClient.UserConfig.Create().SetPlayerID(cookieValue).SetPlayerHP(72).SetPlayerEnergy(3).SetLadder(way).SetCards(initialHand).SetPlayerID(cookieValue).Save(context.Background())
	util.SuccessResp(c, gin.H{"message": "newgame"})
}
