package main

import "fmt"

func main() {
	var s []int
	printAppendSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printAppendSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printAppendSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printAppendSlice(s)
}

func printAppendSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
