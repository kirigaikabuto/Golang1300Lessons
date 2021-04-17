package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func main() {
	var r, cX, cY, x, y string
	//r - радиус окружности
	//cX,cY - координаты центра окружности
	//x,y - координаты какой либо точки
	//находится ли точка x,y внутри окружности или на ее линии
	//евклидово расстояние
	fmt.Print("r=")
	fmt.Scan(&r)
	fmt.Print("cx=")
	fmt.Scan(&cX)
	fmt.Print("cy=")
	fmt.Scan(&cY)
	fmt.Print("x=")
	fmt.Scan(&x)
	fmt.Print("y=")
	fmt.Scan(&y)
	rInt, err := strconv.Atoi(r)
	if err != nil {
		log.Fatal(err)
	}
	cXInt, err := strconv.Atoi(cX)
	if err != nil {
		log.Fatal(err)
	}
	cYInt, err := strconv.Atoi(cY)
	if err != nil {
		log.Fatal(err)
	}
	xInt, err := strconv.Atoi(x)
	if err != nil {
		log.Fatal(err)
	}
	yInt, err := strconv.Atoi(y)
	if err != nil {
		log.Fatal(err)
	}
	part1 := math.Pow(float64(xInt-cXInt), 2.0)
	part2 := math.Pow(float64(yInt-cYInt), 2.0)
	d := math.Sqrt(part1 + part2)
	if d <= float64(rInt) {
		fmt.Println("correct")
	} else {
		fmt.Println("not correct")
	}
}
