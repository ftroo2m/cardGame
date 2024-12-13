package handles

import (
	"cardGame/config"
	"cardGame/ent/userconfig"
	"cardGame/internal/util"
	"context"
	"github.com/gin-gonic/gin"
)

func Card(c *gin.Context) {
	cookieName := "playerID"
	cookieValue, _ := c.Cookie(cookieName)
	usercf, _ := config.SqlClient.UserConfig.Query().Where(userconfig.PlayerID(cookieValue)).First(context.Background())
	util.SuccessResp(c, gin.H{"cards": usercf.Cards})
}
