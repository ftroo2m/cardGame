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

func Rest(c *gin.Context) {
	cookieName := "playerID"
	cookieValue, _ := c.Cookie(cookieName)
	uc, _ := config.SqlClient.UserConfig.Query().Where(userconfig.PlayerID(cookieValue)).First(context.Background())
	if uc.PlayerHP+20 > 72 {
		uc.PlayerHP = 72
	} else {
		uc.PlayerHP += 20
	}
	usercf, _ := config.SqlClient.UserConfig.Query().Where(userconfig.PlayerID(cookieValue)).First(context.Background())
	var nodes []model.Node
	json.Unmarshal([]byte(usercf.Ladder), &nodes)
	for i := 0; i < len(nodes); i++ {
		if nodes[i].Name == "campfire" && nodes[i].Visit != 1 {
			nodes[i].Visit = 1
			break
		}
	}
	jsonWay, _ := json.Marshal(nodes)
	config.SqlClient.UserConfig.Update().Where(userconfig.PlayerID(cookieValue)).SetPlayerHP(uc.PlayerHP).SetLadder(string(jsonWay)).Save(context.Background())
	util.SuccessResp(c, gin.H{"playerHP": uc.PlayerHP})
}
