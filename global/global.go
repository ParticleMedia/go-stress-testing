package global

import (
	"context"

	"golang.org/x/time/rate"
)

var EnableStepSync = false

var ReadyCh = make(chan bool)
var DoneCh = make(chan bool)

func Feed(qps int, concurrency uint64) {
	limiter := rate.NewLimiter(rate.Limit(qps)/rate.Limit(concurrency), 1)
	var i uint64
	for {
		limiter.Wait(context.Background())

		for i = 0; i < concurrency; i++ {
			ReadyCh <- true
		}

		for i = 0; i < concurrency; i++ {
			<-DoneCh
		}
	}
}
