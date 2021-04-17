package main

import "fmt"

func main() {
	var start, end int
	fmt.Scan(&start)
	fmt.Scan(&end)
	for i := start; i <= end; i++ {
		fmt.Println(i)
	}
}
