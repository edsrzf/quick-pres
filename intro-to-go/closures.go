package main

import (
	"fmt"
)

func main() {
	x := 1
	addToX := func(y int) {
		x += y
	}
	addToX(1)
	fmt.Println(x)
}
