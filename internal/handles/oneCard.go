package handles

import (
	"cardGame/config"
	"cardGame/ent/card"
	"cardGame/internal/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

type CardResponse struct {
	Name string `json:"name"`
}

func OneCard(c *gin.Context) {
	var cardname CardResponse
	c.BindJSON(&cardname)
	cd, _ := config.SqlClient.Card.Query().Where(card.Name(cardname.Name)).First(context.Background())
	cd.Image = util.ImageTobase64(cd.Name)
	util.SuccessResp(c, gin.H{"card": cd})
}
