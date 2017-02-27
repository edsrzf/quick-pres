package main

import (
	"fmt"
	"sync"
	"time"
)

// START OMIT
func printWorker(ch chan string, wg *sync.WaitGroup) {
loop:
	for {
		select {
		case s, ok := <-ch:
			if !ok {
				return
			}
			fmt.Println(s)
		case <-time.After(1 * time.Minute):
			fmt.Println("timed out")
			break loop
		}
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
