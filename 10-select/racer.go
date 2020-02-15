package racer

import (
	"net/http"
)

func Racer(urlA, urlB string) (winner string, err error) {
	select {
	case <- ping(urlA):
		return urlA, nil
	case <- ping(urlB):
		return urlB, nil
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
