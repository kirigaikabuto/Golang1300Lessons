package main

import "fmt"

func main() {
	a := 5
	var b *int
	b = &a
	fmt.Println(&b, b, *b)
	fmt.Println(&a, a)
}
