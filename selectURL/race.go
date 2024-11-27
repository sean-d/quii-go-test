package selectURL

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(a, b string) (string, error) {

	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(10 * time.Millisecond):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}

	//startA := time.Now()
	//_, _ = http.Get(a)
	//aDuraction := time.Since(startA)
	//
	//startB := time.Now()
	//_, _ = http.Get(b)
	//bDuraction := time.Since(startB)
	//
	//if aDuraction > bDuraction {
	//	return b
	//}
	//return b
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		_, _ = http.Get(url)
		close(ch)
	}()
	return ch
}
