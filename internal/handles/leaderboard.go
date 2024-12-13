package handles

import (
	"cardGame/config"
	"cardGame/internal/util"
	"context"
	"github.com/gin-gonic/gin"
)

type UserResponse struct {
	PlayerID string `json:"playerID"`
	Counts   int    `json:"counts"`
	// 添加其他你想要返回的字段
}

func LeaderBoard(c *gin.Context) {
	playList, err := config.SqlClient.Leaderboard.Query().All(context.Background())
	if err != nil {
		util.ErrorResp(c, err, 500, true)
		return
	}

	userResponses := make([]UserResponse, len(playList))
	for i, player := range playList {
		userResponses[i] = UserResponse{
			PlayerID: player.PlayerID,
			Counts:   player.Counts,
		}
	}

	util.SuccessResp(c, gin.H{"leaderboard": userResponses})
}
