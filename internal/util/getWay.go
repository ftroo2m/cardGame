package util

import (
	"cardGame/config"
	"cardGame/ent/monster"
	"cardGame/internal/model"
	"context"
	"encoding/json"
	"math/rand"
	"sort"
	"time"
)

func GetWay() string {
	way := make([]model.Node, 8)
	rand.Seed(time.Now().UnixNano())
	id := [7]int{}
	p := 0
	for i := 1; i <= 3; i++ {
		sid := rand.Intn(6) + 3
		for j := 1; j < i; j++ {
			if sid == id[j] {
				p = 1
				break
			}
		}
		if p == 1 {
			i--
			p = 0
		} else {
			id[i] = sid
		}
	}
	for i := 4; i <= 5; i++ {
		sid := rand.Intn(3) + 9
		for j := 4; j < i; j++ {
			if sid == id[j] {
				p = 1
				break
			}
		}
		if p == 1 {
			i--
			p = 0
		} else {
			id[i] = sid
		}
	}
	id[6] = rand.Intn(2) + 1
	sort.Ints(id[:])
	for i := 1; i <= 3; i++ {
		monster, _ := config.SqlClient.Monster.Query().Where(monster.ID(id[i])).First(context.Background())
		way[i-1] = model.Node{
			ID:    i - 1,
			Name:  monster.Name,
			Visit: 0,
		}
	}
	way[3] = model.Node{
		ID:    3,
		Name:  "campfire",
		Visit: 0,
	}
	for i := 4; i <= 5; i++ {
		monster, _ := config.SqlClient.Monster.Query().Where(monster.ID(id[i])).First(context.Background())
		way[i] = model.Node{
			ID:    i,
			Name:  monster.Name,
			Visit: 0,
		}
	}
	way[6] = model.Node{
		ID:    6,
		Name:  "campfire",
		Visit: 0,
	}
	monster, _ := config.SqlClient.Monster.Query().Where(monster.ID(id[6])).First(context.Background())
	way[7] = model.Node{
		ID:    7,
		Name:  monster.Name,
		Visit: 0,
	}
	jsonWay, _ := json.Marshal(way)
	return string(jsonWay)
}
