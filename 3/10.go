package main

import "fmt"

func main() {
	sumi := 5
	fmt.Println("start cycle")
	for i := 0; i < 5; i++ {
		//заходим
		fmt.Println("i=", i)
		fmt.Println("before sumi=",sumi)
		sumi = sumi + 1
		fmt.Println("after sumi=",sumi)
		fmt.Println("")
		//выходим
	}
	fmt.Println("end cycle")
	//в начале
	//i=0; sumi = 5; sumi = 6
	//i=1; sumi = 6; sumi = 7
	//i=2; sumi = 7; sumi = 8
	//i=3; sumi = 8;
	//i=4; sumi = 9;
	//i=5; sumi = 10; i<5 -> false
	fmt.Println(sumi)
}
