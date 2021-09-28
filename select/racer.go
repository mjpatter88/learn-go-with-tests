package racer

import (
	"net/http"
)

func Racer(a string, b string) (winner string) {
	// select blocks on reading from multiple channels.
	// The first case that receives a value gets executed.
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}

// Return a channel that is closed when the request completes
func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}
