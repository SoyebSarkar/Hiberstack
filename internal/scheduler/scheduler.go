package scheduler

import (
	"time"
)

func Start(fn func(), interval time.Duration) {
	go func() {
		for {
			fn()
			time.Sleep(interval)
		}
	}()
}
