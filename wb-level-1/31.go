package main

import (
	"fmt"
	"math"
)

// Написать программу нахождения расстояния между 2 точками, которые представление
// в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

func main() {
	p1 := NewPoint(4.7, 9.5)
	p2 := NewPoint(1.7, 5.5)
	fmt.Println(Distance(p1, p2))
}

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func Distance(p1, p2 *Point) float64 {
	return math.Sqrt((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y))
}
