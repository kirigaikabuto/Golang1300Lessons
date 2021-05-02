package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//strconv.Itoa-> numbers to string
	//strconv.Atoi-> string to numbers
	n := 2500
	primeNumbers := []string{}
	for i := 1; i < n; i++ {
		counter := 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				counter += 1
			}
		}
		if counter == 2 {
			primeNumbers = append(primeNumbers, strconv.Itoa(i))
		}
	}
	s := strings.Join(primeNumbers, "")
	fmt.Println(len(s))
	arr := []int{2, 5, 6, 8, 12}
	for _, v := range arr {
		fmt.Print(string(s[v-1]))
	}
}
