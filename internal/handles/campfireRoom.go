package handles

import (
	"cardGame/config"
	"cardGame/ent/userconfig"
	"cardGame/internal/util"
	"context"
	"github.com/gin-gonic/gin"
)

func Rest(c *gin.Context) {
	cookieName := "playerID"
	cookieValue, _ := c.Cookie(cookieName)
	uc, _ := config.SqlClient.UserConfig.Query().Where(userconfig.PlayerID(cookieValue)).First(context.Background())
	if uc.PlayerHP+20 > 72 {
		uc.PlayerHP = 72
	} else {
		uc.PlayerHP += 20
	}
	config.SqlClient.UserConfig.Update().Where(userconfig.PlayerID(cookieValue)).SetPlayerHP(uc.PlayerHP).Save(context.Background())
	util.SuccessResp(c, gin.H{"playerHP": uc.PlayerHP})
}
