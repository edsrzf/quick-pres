package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)

	v.X = 4
	fmt.Println(v.X, v.Y)

	v2 := &Vertex{X: 1}
	fmt.Println(v2.X, v2.Y)
}
