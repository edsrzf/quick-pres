package main

import "fmt"

type MyFloat float64

func (f MyFloat) Abs() MyFloat {
	if f < 0 {
		return -f
	}
	return f
}

func main() {
	f := MyFloat(-3.14159)
	fmt.Println(f.Abs())
}
