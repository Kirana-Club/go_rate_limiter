package main

import (
	"time"

	"github.com/Kirana-Club/go_rate_limiter"
)

func main() {
	r, err := ratelimiter.NewMaxConcurrencyRateLimiter(&ratelimiter.Config{
		Limit:            4,
		TokenResetsAfter: 10 * time.Second,
	})

	if err != nil {
		panic(err)
	}

	ratelimiter.DoWork(r, 10)
}
