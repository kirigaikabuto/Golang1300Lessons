package main

import (
	"fmt"
	"github.com/kirigaikabuto/Golang1300Lessons/19/users"
)

func main() {
	mongoUsersStore, err := users.NewUsersStore(users.MongoConfig{
		Host:           "localhost",
		Port:           "27017",
		Database:       "lesson19",
		CollectionName: "users",
	})
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(mongoUsersStore.Create(&users.User{
		Username: "kirito",
		Password: "kirito",
	}))
}
