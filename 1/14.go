package main

import "fmt"

func main() {
	var name string
	var age int
	fmt.Print("Name:")
	fmt.Scan(&name)
	fmt.Print("Age:")
	fmt.Scan(&age)
	year := 2021 - age
	out := fmt.Sprintf("Hello %s you was born in %d", name, year)
	fmt.Println(out)

}
