package main

//If the top-level type is just a type name, you can omit it from the elements of the literal.
import "fmt"

type Vertex1 struct {
	Lat, Long float64
}

var m2 = map[string]Vertex1{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m2)
}
