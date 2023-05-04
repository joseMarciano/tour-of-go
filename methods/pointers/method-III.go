package main

import (
	"fmt"
	"math"
)

type VertexPointer struct {
	X, Y float64
}

func (v VertexPointer) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *VertexPointer) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := VertexPointer{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
