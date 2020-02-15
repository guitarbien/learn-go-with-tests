package racer

import (
	"net/http"
)

func Racer(urlA, urlB string) (winner string) {
	select {
	case <- ping(urlA):
		return urlA
	case <- ping(urlB):
		return urlB
	}
}

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()

	return ch
}
