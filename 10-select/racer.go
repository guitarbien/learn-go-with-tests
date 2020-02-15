package racer

import (
	"net/http"
	"time"
)

func Racer(urlA, urlB string) (winner string) {
	urlADuration := measureResponseTime(urlA)
	urlBDuration := measureResponseTime(urlB)

	if urlADuration < urlBDuration {
		return urlA
	}

	return urlB
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)

	return time.Since(start)
}
