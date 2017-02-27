package main

import (
	"fmt"
	"sync"
)

// START OMIT
func printWorker(ch chan string, wg *sync.WaitGroup) {
	for s := range ch { // range over channel
		fmt.Println(s)
	}
	wg.Done()
}

func main() {
	const workers = 2
	ch := make(chan string, workers) // buffered channel
	var wg sync.WaitGroup
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go printWorker(ch, &wg)
	}
	for _, s := range []string{"goroutines", "are", "fun"} {
		ch <- s
	}
	close(ch) // close terminates range loop
	wg.Wait()
}

// END OMIT
