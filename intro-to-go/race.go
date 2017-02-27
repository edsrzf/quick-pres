package main

import "fmt"

// START OMIT
func main() {
	const (
		neverUsed = "short garbage"
		long      = "a long string"
	)
	short := neverUsed[:5]
	s := short
	go func() {
		for {
			s = short
			s = long
		}
	}()
	for {
		if s == neverUsed {
			fmt.Println("Race! s == ", s)
			break
		}
	}
}

// END OMIT
