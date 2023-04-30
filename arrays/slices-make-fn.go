package main

import "fmt"

func main() {
	a := make([]int, 5)
	slicePrint("a", a)

	b := make([]int, 0, 5)
	slicePrint("b", b)

	c := b[:2]
	slicePrint("c", c)

	d := c[2:5]
	slicePrint("d", d)
}

func slicePrint(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
