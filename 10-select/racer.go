package racer

import (
	"net/http"
	"time"
)

func Racer(urlA, urlB string) (winner string) {
	startA := time.Now()
	http.Get(urlA)
	urlADuration := time.Since(startA)

	startB := time.Now()
	http.Get(urlB)
	urlBDuration := time.Since(startB)

	if urlADuration < urlBDuration {
		return urlA
	}

	return urlB
}
