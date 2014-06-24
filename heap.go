// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import (
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"
)

// START1 OMIT
type Heap []int

func (Heap) Generate(rand *rand.Rand, size int) reflect.Value {
	var h Heap
	for i := 0; i < size; i++ {
		switch rand.Intn(4) {
		case 0, 1:
			h.Push(rand.Int())
		case 2:
			if len(h) > 0 {
				h.Pop()
			}
		case 3:
			if len(h) > 0 {
				h.Remove(rand.Intn(len(h)))
			}
		}
	}
	return reflect.ValueOf(h)
}
// END1 OMIT

// START2 OMIT
func TestHeap(t *testing.T) {
	f := func(h Heap) bool {
		if len(h) <= 1 {
			return true
		}
		last := h.Pop()
		for len(h) > 0 {
			popped := h.Pop()
			if popped < last {
				return false
			}
			last = popped
		}
		return true
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
// END2 OMIT

func (h *Heap) Push(x int) {
	n := len(*h)
	*h = append(*h, x)
	h.up(n)
}

func (h *Heap) Pop() int {
	n := len(*h) - 1
	(*h)[0], (*h)[n] = (*h)[n], (*h)[0]
	h.down(0, n)
	el := (*h)[n]
	*h = (*h)[:n]
	return el
}

func (h *Heap) Remove(i int) int {
	n := len(*h) - 1
	if n != i {
		(*h)[i], (*h)[n] = (*h)[n], (*h)[i]
		h.down(i, n)
		h.up(i)
	}
	el := (*h)[n]
	*h = (*h)[:n]
	return el
}

func (h *Heap) up(j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || (*h)[j] >= (*h)[i] {
			break
		}
		(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
		j = i
	}
}

func (h *Heap) down(i, n int) {
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && (*h)[j1] >= (*h)[j2] {
			j = j2 // = 2*i + 2  // right child
		}
		if (*h)[j] >= (*h)[i] {
			break
		}
		(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
		i = j
	}
}

func main() {
	match := func(string, string) (bool, error) { return true, nil }
	tests := []testing.InternalTest{{"TestHeap", TestHeap}}
	testing.Main(match, tests, nil, nil)
}
