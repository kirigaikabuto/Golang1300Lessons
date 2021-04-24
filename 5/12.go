package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	n := 5
	a := []int{}
	for i := 0; i < n; i++ {
		k := rand.Intn(100)
		a = append(a, k)
	}
	fmt.Println(a)
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if a[i] < a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	fmt.Println(a)
}
