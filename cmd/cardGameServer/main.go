package main

import (
	"cardGame/config"
	"cardGame/internal/server"
	"github.com/gin-gonic/gin"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	config.Init()
	r := gin.Default()
	server.Init(r)
	r.Run("127.0.0.1:8080")
	//User, err := config.SqlClient.
	//	User.Query().
	//	Where(user.Username("admin")).
	//	First(context.Background())
	//if err != nil {
	//	panic(err)
	//}
	//println(User.Password)

}
