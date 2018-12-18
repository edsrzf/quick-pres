package main

import (
	"fmt"
	"os"
)

// START OMIT
func main() {
	// func Open(name string) (*File, error)
	f, err := os.Open("file-not-found")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("file not found!")
			return
		}
		if perr, ok := err.(*os.PathError); ok {
			fmt.Println("path error at", perr.Path)
			return
		}
		fmt.Println("other error", err)
	}
	_ = f
}

// END OMIT
