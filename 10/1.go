package main

import "github.com/kirigaikabuto/Golang1300Lessons/10/users"

func main() {
	user1 := users.User{
		Id:       "1",
		Username: "kirito",
		Password: "12123123",
		Marks:    []int{1, 2, 3, 4},
		Address: users.AddressModel{
			District:   "13",
			Street:     "asdsadas",
			HomeNumber: "34",
		},
	}
	marks2 := []int{3, 4, 5, 3}
	user2 := users.User{
		Id:       "2",
		Username: "asdsada",
		Password: "13123123",
		Marks:    marks2,
		Address: users.AddressModel{
			District:   "1",
			Street:     "33333",
			HomeNumber: "123",
		},
	}
	usersData := []users.User{user1, user2}
	for _, v := range usersData {
		v.ShowInfo()
	}
}
