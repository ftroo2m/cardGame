package handles

import "github.com/gin-gonic/gin"

func Map(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})

}
