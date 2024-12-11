package handles

import "github.com/gin-gonic/gin"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

}

func Logout(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func Register(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
