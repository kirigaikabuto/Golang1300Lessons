package main

import "fmt"

func main() {
	var n, y, p float64
	fmt.Scan(&n)
	fmt.Scan(&y)
	fmt.Scan(&p)
	m := (n * p * (1 + p) * y) / (12 * ((1+p)*y - 1))
	s := (m * 12) * y
	fmt.Println(m, s)
}
