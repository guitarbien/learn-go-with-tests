package racer

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(urlA, urlB string) (winner string, err error) {
	select {
	case <- ping(urlA):
		return urlA, nil
	case <- ping(urlB):
		return urlB, nil
	case <- time.After(10 * time.Second):
		return "", fmt.Errorf("timed out waiting for %s and %s", urlA, urlB)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}
