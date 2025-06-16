package config

import (
	"cardGame/ent"
	"context"
	"entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq"
	"log"
	"time"
)

var DataBase = "postgres://user:password@ip:port/cardgame?sslmode=disable"

var SqlClient *ent.Client

var MonsterList = map[string]ent.Monster{}

var CardList = map[string]ent.Card{}

func Init() {
	var err error
	SqlClient, err = createDatabase()
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}
	log.Println("Database connection initialized successfully.")

	getCard()

	getMonster()
}

func createDatabase() (*ent.Client, error) {
	drv, err := sql.Open("postgres", DataBase)
	if err != nil {
		return nil, err
	}
	db := drv.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	return ent.NewClient(ent.Driver(drv)), nil
}

func getCard() {
	Cardlist, _ := SqlClient.Card.Query().All(context.Background())
	for _, card := range Cardlist {
		CardList[card.Name] = *card
	}
}

func getMonster() {
	Monsterlist, _ := SqlClient.Monster.Query().All(context.Background())
	for _, monster := range Monsterlist {
		MonsterList[monster.Name] = *monster
	}
}
