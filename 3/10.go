package main

import "fmt"

func main() {
	sumi := 5
	for i := 0; i < 5; i++ {
		sumi = sumi + 1
	}
	//в начале
	//i=0; sumi = 5;
	//i=1; sumi = 6;
	//i=2; sumi = 7;
	//i=3; sumi = 8;
	//i=4; sumi = 9;
	//i=5; sumi = 10; i<5 -> false
	fmt.Println(sumi)
}
