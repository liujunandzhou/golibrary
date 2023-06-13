package main

import (
	"fmt"
	"log"
	"time"

	"github.com/throttled/throttled/store/memstore"
	"github.com/throttled/throttled/v2"
)

func main() {

	store, err := memstore.New(65536)
	if err != nil {
		log.Printf("memstore.New failed,err=%v\n", err)
		return
	}

	quota := throttled.RateQuota{
		MaxRate:  throttled.PerMin(20),
		MaxBurst: 10,
	}

	rateLimiter, err := throttled.NewGCRARateLimiter(store, quota)
	if err != nil {
		log.Printf("throtted.NewGCRARateLimiter failed,err=%v\n", err)
		return
	}

	key := "tester"

	for {

		status, res, err := rateLimiter.RateLimit(key, 1)
		if err != nil {
			log.Printf("rateLimiter.RateLimit failed,err=%v\n", err)
			continue
		}

		fmt.Printf("status=%v res=%+v\n", status, res)

		time.Sleep(time.Second * 3)
	}
}
