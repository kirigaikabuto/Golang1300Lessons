package main

import (
	"fmt"
	"strconv"
)

func main() {
	s1 := "asd"
	s2 := "4"
	n1, e1 := strconv.Atoi(s1)
	fmt.Println("err1=", e1)
	n2, e2 := strconv.Atoi(s2)
	fmt.Println("err2=", e2)
	fmt.Println("n1=", n1)
	fmt.Println("n2=", n2)
	c := n1 + n2
	fmt.Println("c=", c)
	s3 := strconv.Itoa(c)
	fmt.Println(s1 + s2 + s3)
}
