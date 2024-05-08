package main

import (
	"log"
	"math"
)

type Point struct {
	a float64
	b float64
	x float64
	y float64
}

func NewPoint(x, y, a, b float64) Point {
	if math.Pow(y, 2) != math.Pow(x, 3)+(a*x)+b {
		log.Panicf("%+v, %+v is not on the curve", x, y)
	}
	return Point{a, b, x, y}
}
