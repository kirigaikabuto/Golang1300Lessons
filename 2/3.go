package main

import (
	"fmt"
	"reflect"
)

func main() {
	k := 3.15
	fmt.Println(reflect.TypeOf(k))
	fmt.Println(k)
	l := int(k)
	fmt.Println(reflect.TypeOf(l))
	fmt.Println(l)
}
