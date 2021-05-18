package main

import (
	"fmt"
	"reflect"
)

func showValue(a interface{}) {
	fmt.Println(a, reflect.TypeOf(a))
}

func main() {
	s := "asdsdadasd"
	showValue(s)
	a := 3
	showValue(a)
}
