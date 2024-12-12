package handles

import "github.com/gin-gonic/gin"

func Rest(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})

}

func Up(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func Down(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}