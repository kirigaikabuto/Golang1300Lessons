package main

import "fmt"

func setSum(a, b int, c *int) {
	*c = a + b
}

func main() {
	var otvet int
	setSum(3, 4, &otvet)
	fmt.Println(otvet)
}
