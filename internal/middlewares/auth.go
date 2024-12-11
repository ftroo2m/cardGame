package middlewares

import "github.com/gin-gonic/gin"

func Auth(c *gin.Context) {
	c.Next()
}
