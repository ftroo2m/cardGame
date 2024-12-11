package config

import (
	"cardGame/ent"
	"cardGame/internal/model"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql" // 加载 MySQL 驱动
	"log"
	"time"
)

var GameManager = model.NewGameManager()

var DataBase = "root:0987654321@tcp(127.0.0.1:3306)/cardgame?parseTime=True"

var SqlClient *ent.Client

func Init() {
	var err error
	SqlClient, err = CreateDatabase()
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}
	log.Println("Database connection initialized successfully.")
}

func CreateDatabase() (*ent.Client, error) {
	drv, err := sql.Open(dialect.MySQL, DataBase)
	if err != nil {
		return nil, err
	}
	db := drv.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	return ent.NewClient(ent.Driver(drv)), nil
}
