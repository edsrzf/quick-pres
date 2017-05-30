package main

import (
	"net/http"
)

func main() {
	err := http.ListenAndServe(":8080", http.HandlerFunc(helloHandler))
	if err != nil {
		panic(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
