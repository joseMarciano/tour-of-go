package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13, 14, 15, 16}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[2:4]
	printSlice(s)

	// Extend its length.
	s = s[:7]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
