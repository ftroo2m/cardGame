package handles

import (
	"cardGame/config"
	"cardGame/ent/user"
	"cardGame/internal/model"
	"cardGame/internal/util"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Account struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var req Account
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		util.ErrorStrResp(c, "JSON invalid", 400, true)
		return
	}
	User, err := config.SqlClient.User.
		Query().
		Where(user.Username(req.Username)).
		First(context.Background())
	if err != nil {
		util.ErrorResp(c, err, 400, true)
	}

	if User.Password != req.Password {
		util.ErrorStrResp(c, "Password error", 401, true)
	} else {
		c.SetCookie("playerID", User.Username, 86400, "/", "localhost", false, false)
		util.SuccessResp(c)
	}
}

func Logout(c *gin.Context) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "playerID",      // cookie的名称
		Value:    "",              // cookie的值设置为空
		Expires:  time.Unix(0, 0), // 设置过期时间为1970年1月1日
		HttpOnly: false,           // 设置HttpOnly标志，提高安全性
		Path:     "/",             // 设置cookie的有效路径
	})
	util.SuccessResp(c)
}

func Register(c *gin.Context) {
	var req Account
	err := c.BindJSON(&req)
	if err != nil {
		util.ErrorResp(c, err, http.StatusBadRequest, true)
	}
	u, _ := config.SqlClient.
		User.Query().Where(user.Username(req.Username)).First(context.Background())
	if u != nil {
		util.ErrorResp(c, err, 301, true)
	}
	way := model.GetWay()
	initialHand := []string{"袭击", "袭击", "袭击", "袭击", "袭击", "防御", "防御", "防御", "防御", "猛击"}
	_, err = config.SqlClient.UserConfig.Create().SetPlayerID(req.Username).SetPlayerHP(72).SetPlayerEnergy(3).SetLadder(way).SetCards(initialHand).Save(context.Background())
	if err != nil {
		panic(err)
		util.ErrorResp(c, err, 401, true)
	}
	_, err = config.SqlClient.
		User.Create().SetUsername(req.Username).SetPassword(req.Password).Save(context.Background())
	if err != nil {
		panic(err)
		util.ErrorResp(c, err, 401, true)
	}
	ur, _ := config.SqlClient.
		User.Query().Where(user.Username(req.Username)).First(context.Background())
	go util.NewPlayer(ur.ID + 5)
	imageBase64 := util.ImageToBase64Player(ur.ID)
	c.SetCookie("playerID", ur.Username, 86400, "/", "localhost", false, false)
	util.SuccessResp(c, gin.H{
		"image": imageBase64,
	})
}
