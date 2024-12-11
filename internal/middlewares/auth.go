package middlewares

import (
	"cardGame/config"
	"cardGame/ent/user"
	"cardGame/internal/util"
	"context"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	cookieName := "playerID"
	cookieValue, err := c.Cookie(cookieName)
	if err != nil {
		util.ErrorResp(c, err, 402, true)
		c.Abort()
	} else {
		_, err := config.SqlClient.User.Query().Where(user.Username(cookieValue)).First(context.Background())
		if err != nil {
			util.ErrorResp(c, err, 402, true)
			c.Abort()
		} else {
			c.Next()
		}
	}

	c.Next()
}
