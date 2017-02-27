package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println(s)

	primeSlice := []int{2, 3, 5, 7, 11, 13}
	primeSlice = append(primeSlice, 17)
	fmt.Println(primeSlice)

	zeroSlice := make([]int, 6)
	fmt.Println(zeroSlice)
}
