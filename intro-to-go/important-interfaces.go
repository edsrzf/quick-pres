package main

import (
	"net/http"
)

// START OMIT
// error
type error interface {
	Error() string
}

// io.Reader
type Reader interface {
	Read([]byte) (int, error)
}

// io.Writer
type Writer interface {
	Write([]byte) (int, error)
}

// http.Handler
type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

// END OMIT

func main() {
}
