package db

import (
	"fmt"
	"math/rand"
	"paw-me-back/internal/model"

	"github.com/google/uuid"
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
	users := generateUsers(100)
	db.CreateInBatches(&users, len(users))

	groups := generateGroups(50, 10, users)
	db.CreateInBatches(&groups, len(groups))
}

func generateUsers(count int) []model.User {
	users := make([]model.User, count)

	for i := 0; i < count; i++ {
		users[i] = model.User{
			Username:     usernamesRepo[i%len(usernamesRepo)] + fmt.Sprintf("%d", i),
			Email:        usernamesRepo[i%len(usernamesRepo)] + fmt.Sprintf("%d", i) + "@gmail.com",
			SuperTokenID: uuid.New().String(),
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
