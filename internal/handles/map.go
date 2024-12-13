package handles

import (
	"cardGame/config"
	"cardGame/ent/userconfig"
	"cardGame/internal/model"
	"cardGame/internal/util"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func Map(c *gin.Context) {
	cookieName := "playerID"
	cookieValue, _ := c.Cookie(cookieName)
	usercf, _ := config.SqlClient.UserConfig.Query().Where(userconfig.PlayerID(cookieValue)).First(context.Background())
	var nodes []model.Node
	json.Unmarshal([]byte(usercf.Ladder), &nodes)
	util.SuccessResp(c, gin.H{"nodes": nodes})
}
