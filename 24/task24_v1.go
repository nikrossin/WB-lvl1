package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) *Point {
	return &Point{x, y}
}

func (p1 *Point) Distance(p2 *Point) float64 {
	return math.Sqrt((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y))
}

func main() {
	p1 := NewPoint(0, 10)
	p2 := NewPoint(0, 0)
	fmt.Println(p1.Distance(p2))
}
