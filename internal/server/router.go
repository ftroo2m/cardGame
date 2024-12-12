package server

import (
	"cardGame/internal/handles"
	"cardGame/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func Init(e *gin.Engine) {
	api := e.Group("/api")

	api.Any("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api.POST("/auth/login", handles.Login)
	api.POST("/auth/logout", handles.Logout)
	api.POST("/auth/register", handles.Register)

	g := api.Group("/game", middlewares.Auth)
	g.POST("/leaderboard", handles.LeaderBoard)
	g.POST("/start/game", handles.StartGame)
	g.POST("/map", handles.Map)
	g.POST("/c/rest", handles.Rest)
	g.POST("/c/up", handles.Up)
	g.POST("/c/down", handles.Down)
	g.POST("/win", handles.Win)
	//g.POST("/archive", handles.Archive)
}
