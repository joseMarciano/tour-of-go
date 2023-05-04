package main

import (
	"fmt"
	"math"
)

type VertexIndirectionV struct {
	X, Y float64
}

func (v VertexIndirectionV) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v VertexIndirectionV) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := VertexIndirectionV{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &VertexIndirectionV{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}
