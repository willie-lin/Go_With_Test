package recer

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {

	startA := time.Now()
	http.Get(a)
	aDuration := time.Since(startA)

	startB := time.Now()
	http.Get(b)
	bDuraction := time.Since(startB)

	if aDuration < bDuraction {
		return a
	}
	return b
}
