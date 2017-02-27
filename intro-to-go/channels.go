package main

import (
	"fmt"
	"time"
)

// START OMIT
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	ch := make(chan bool)
	go func() {
		say("world")
		ch <- true
	}()
	say("hello")
	<-ch
}

// END OMIT
