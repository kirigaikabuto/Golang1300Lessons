package main

import (
	"math/rand"
	"time"
)

func getNewArray(n int) []int {
	arr := []int{}
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < n; i++ {
		k := rand.Intn(100)
		arr = append(arr, k)
	}
	return arr
}
//функция для посчета суммы элементов в массиве и возвращать эту сумму
func main() {
	a := getNewArray(10)
	b := getNewArray(20)
	allArr := [][]int{a, b}
}
