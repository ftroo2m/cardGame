package handles

import (
	"cardGame/config"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type StartGameRequest struct {
	Monster string `json:"monster"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// 允许所有请求，实际项目中可以根据需求限制跨域
		return true
	},
}

func StartGame(c *gin.Context) {
	var req StartGameRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	cookieName := "playerID"
	playerID, err := c.Cookie(cookieName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "playerID is required"})
		return
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Failed to upgrade to WebSocket:", err)
		return
	}
	config.GameManager.CreateGame(playerID, req.Monster, conn)
}
