package main

import (
	"fmt"
	"github.com/kirigaikabuto/Golang1300Lessons/9/persons"
)

func main() {
	u1 := persons.User{
		Name:     "Yerassyl",
		Username: "kirigaiakabuto",
		Password: "123212312",
		Age:      23,
	}
	u1.ShowName()
	u1.ShowUsernameAndPassword()
	fmt.Println(u1.YearOfBirth)
	u1.SetYearOfBirth()
	fmt.Println(u1.YearOfBirth)
}
