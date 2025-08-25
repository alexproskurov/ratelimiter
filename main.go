package main

import (
	"fmt"
	"time"

	"gthub.com/alexproskurov/ratelimiter/ratelimiter"
)

func main() {
	rl := ratelimiter.NewLeakyBucket(5, 200*time.Microsecond)
	for i := range 10 {
		if rl.Allow() {
			fmt.Printf("request %d allowed", i)
		} else {
			fmt.Printf("Request %d rejected", i)
		}
		time.Sleep(100 * time.Millisecond)
	}
}
