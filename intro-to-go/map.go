package main

import "fmt"

func main() {
	lookup := map[string]int{"answer": 42}
	fmt.Println(lookup["answer"])

	lookup["answer"] = 48
	fmt.Println(lookup["answer"])

	delete(lookup, "answer")
	fmt.Println(lookup["answer"])

	lookup2 := make(map[string]int, 50)
	fmt.Println(lookup2["answer"])
}
