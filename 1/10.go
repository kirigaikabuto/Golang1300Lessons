package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int
	fmt.Println(&a)
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a))
}
