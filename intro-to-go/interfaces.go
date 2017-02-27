package main

import "fmt"

// START OMIT
type Printer interface {
	Print()
}

func printAThing(p Printer) { p.Print() }

type MyFloat float64

func (f MyFloat) Print() { fmt.Println(f) }

type MyString string

func (s MyString) Print() { fmt.Println(s) }

func main() {
	printAThing(MyFloat(-3.14159))
	printAThing(MyString("hi"))
}

// END OMIT
