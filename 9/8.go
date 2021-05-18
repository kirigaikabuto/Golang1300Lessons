package main

import (
	"fmt"
	"github.com/kirigaikabuto/Golang1300Lessons/9/persons"
)

func main() {
	c := persons.GetSum(3, 4)
	fmt.Println(c)
	d := persons.GetMultiplication(3, 4)
	fmt.Println(d)
}
