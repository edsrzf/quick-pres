package main

import (
	"reflect"
)

// START OMIT
// A Generator can generate random values of its own type.
type Generator interface {
	// Generate returns a random instance of the type on which it is a
	// method using the size as a size hint.
	Generate(rand *rand.Rand, size int) reflect.Value
}

// END OMIT

func main() {}
