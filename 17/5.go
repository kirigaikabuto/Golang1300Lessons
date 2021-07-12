package main

import (
	"fmt"
)

var (
	cha    = make(chan int)
	global = make(chan int)
)

func Set() {
	fmt.Println("set run")
	cha <- 3
}

func Get() {
	fmt.Println("get run")
	fmt.Println(<-cha)
	cha <- 4
	fmt.Println("end of get")
	global <- 5
}

func main() {
	go Set()
	go Get()
	<-global
}
