package main

import "fmt"

/**
The select statement lets a goroutine wait on multiple communication operations.

A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.
*/

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		fmt.Println("Step 1")
		select {
		case c <- x:
			x, y = y, x+y
			fmt.Printf("Step 2 x: %v  y: %v \n", x, y)
		case <-quit:
			fmt.Println("Step 3")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("Step 5: %v \r\n", <-c)
		}
		fmt.Println("Step 4")
		quit <- 0
	}()
	fibonacci(c, quit)
}
