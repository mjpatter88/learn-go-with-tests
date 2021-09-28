package racer

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(a string, b string) (winner string, err error) {
	// select blocks on reading from multiple channels.
	// The first case that receives a value gets executed.
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(10 * time.Second):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
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
