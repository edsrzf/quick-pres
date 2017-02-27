package main

import (
	"fmt"
	"sync"
	"time"
)

// START OMIT
type SafeCounter struct {
	v map[string]int
	sync.Mutex
}

func (c *SafeCounter) Inc(key string) {
	c.Lock()
	c.v[key]++
	c.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.Lock()
	defer c.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: map[string]int{}}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}

// END OMIT
