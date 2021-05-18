package main

import "github.com/kirigaikabuto/Golang1300Lessons/9/students"

func main() {
	st1 := students.Student{
		Username:  "kirito",
		Password:  "1234",
		FirstName: "Yerassyl",
		LastName:  "Tleugazy",
		Marks:     []int{3, 4, 5, 1, 2, 4, 5, 5, 5},
	}
	st1.ShowInfo()
}
