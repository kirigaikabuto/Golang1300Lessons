package main

import "fmt"

func main() {
	a := 6
	var b *int
	b = &a
	*b = 100
	fmt.Println(a)
	fmt.Println(*b)
}
