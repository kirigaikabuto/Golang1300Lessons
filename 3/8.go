package main

import "fmt"

func main() {
	for counter := 0; counter < 10; counter++ {
		if counter < 5 {
			fmt.Println("Hello")
		} else {
			fmt.Println("World")
		}
	}
}
