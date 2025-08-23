package db

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"paw-me-back/internal/model"

	"gorm.io/gorm"
)

var usernamesRepo = []string{
	"alice", "bob", "charlie", "dave", "eve", "frank", "grace", "heidi",
	"ivan", "judy", "karl", "laura", "mallory", "nina", "oscar", "peggy",
	"quinn", "rachel", "steve", "trent", "ursula", "victor", "wendy", "xander",
	"yvonne", "zack", "amber", "brian", "carol", "doug", "eric", "fiona",
	"george", "hannah", "ian", "jessica", "kevin", "lisa", "mike", "natalie",
	"oliver", "peter", "queen", "ron", "susan", "tim", "uma", "vicky",
	"walter", "xenia", "yasmin", "zoe",
}

func Seed(db *gorm.DB) {
	ctx := context.Background()
	users := generateUsers(100)

	for _, user := range users {
		err := gorm.G[model.User](db).Create(ctx, &user)
		if err != nil {
			log.Println("Error creating user: ", user, err)
			return
		}
	}

	groups := generateGroups(50, 10, users)

	for _, group := range groups {
		err := gorm.G[model.Group](db).Create(ctx, group)
		if err != nil {
			log.Println("Error creating group: ", group, err)
			return
		}
	}

}

func generateUsers(count int) []model.User {
	users := make([]model.User, count)

	for i := 0; i < count; i++ {
		users[i] = model.User{
			Username: usernamesRepo[i%len(usernamesRepo)] + fmt.Sprintf("%d", i),
			Email:    usernamesRepo[i%len(usernamesRepo)] + fmt.Sprintf("%d", i) + "@gmail.com",
		}
	}

	return users
}

func generateGroups(countGroups int, countMembers int, users []model.User) []*model.Group {
	groups := make([]*model.Group, countGroups)

	for i := 0; i < countGroups; i++ {
		owner := users[rand.Intn(len(users))]

		groups[i] = &model.Group{
			Name:  fmt.Sprintf("group-%02d", i+1),
			Owner: owner,
		}

		groups[i].Members = append(groups[i].Members, owner)

		for j := 0; j < countMembers; j++ {
			member := users[rand.Intn(len(users))]
			groups[i].Members = append(groups[i].Members, member)
		}
	}

	return groups
}
