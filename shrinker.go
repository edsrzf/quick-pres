package main

import (
	"reflect"
)

// START1 OMIT
// A Shrinker can produce simpler values with its same type.
type Shrinker interface {
	// Shrink sends simpler values of its type to out. Shrink must stop
	// sending when it receives a value on stop. To avoid possible deadlocks,
	// all sends after the first must be done in a select which also tries to
	// receive from stop.
	Shrink(out chan<- reflect.Value, stop <-chan bool)
}
// END1 OMIT

// START2 OMIT
type Uint uint

func (u Uint) Shrink(out chan<- reflect.Value, stop <-chan bool) {
	out <- reflect.ValueOf(Uint(0))
	i := u / 2
	for i > 0 {
		select {
		case out <- reflect.ValueOf(u - i):
		case <- stop:
			return
		}
		i /= 2
	}
}
// END2 OMIT

// START3 OMIT
func TryShrink(out chan<- reflect.Value, stop <-chan bool, vals ...interface{}) bool {
	for _, val := range vals {
		select {
		case out <- reflect.ValueOf(val):
		case <- stop:
			return true
		}
	}
	return false
}
// END3 OMIT

func main() {}
