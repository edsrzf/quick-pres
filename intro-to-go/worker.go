package main

import "fmt"

// START OMIT
func printWorker(ch chan string, done chan bool) {
	for s := range ch { // range over channel
		fmt.Println(s)
	}
	done <- true
}

func main() {
	const workers = 2
	ch := make(chan string, workers) // buffered channel
	done := make(chan bool)
	for i := 0; i < workers; i++ {
		go printWorker(ch, done)
	}
	for _, s := range []string{"goroutines", "are", "fun"} {
		ch <- s
	}
	close(ch) // close terminates range loop
	for i := 0; i < workers; i++ {
		<-done // wait until workers have finished
	}
}

// END OMIT
