package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 7; i++ {
		go factorial(i)
	}
	fmt.Scanln()
	fmt.Println("end of main")
}
func factorial(n int) {
	if n < 1 {
		fmt.Println("Unvalid input number")
		return
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	fmt.Println("from goroutine=", n, n, "-", result)
}
