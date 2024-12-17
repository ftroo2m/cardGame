package server

import (
	"cardGame/internal/handles"
	"cardGame/internal/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init(e *gin.Engine) {
	e.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // 允许的前端地址
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	api := e.Group("/api")

	api.Any("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//api.Any("/test", func(c *gin.Context) {
	//	Cards, err := config.SqlClient.Card.Query().All(c)
	//	if err != nil {
	//		c.JSON(200, gin.H{
	//			"message": "error",
	//		})
	//	}
	//	cardBase64 := []string{}
	//	for _, card := range Cards {
	//		cardBase64 = append(cardBase64, util.ImageTobase64(card.Name))
	//	}
	//	c.JSON(200, gin.H{
	//		"message": cardBase64,
	//	})
	//})

	api.POST("/auth/login", handles.Login)
	api.POST("/auth/logout", handles.Logout)
	api.POST("/auth/register", handles.Register)

	g := api.Group("/game", middlewares.Auth)
	g.POST("/leaderboard", handles.LeaderBoard)
	g.GET("/start/game", handles.StartGame)
	g.POST("/map", handles.Map)
	g.POST("/c/rest", handles.Rest)
	g.POST("/newgame", handles.NewGame)
	g.POST("/card", handles.Card)
	g.POST("/onecard", handles.OneCard)
	//g.POST("/archive", handles.Archive)
}
