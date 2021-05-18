package main

import "fmt"

func addElement(arr *[]int, inset int) {
	*arr = append(*arr, inset)
}

func main() {
	a := []int{1, 2, 3, 4}
	fmt.Println(a)
	addElement(&a, 100)
	fmt.Println(a)
}
