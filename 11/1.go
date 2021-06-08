package main

import "fmt"

func main() {
	//d := []int{1, 2, 3, 4}
	//values 1 2 3 4 :int
	//index  0 1 2 3 :int
	di := map[string]int{
		"one": 123,
		"two": 345,
	}
	di["three"] = 323
	fmt.Println(di["two"])
	fmt.Println(di["three"])
	for i, v := range di {
		fmt.Println(i, v)
	}
}
