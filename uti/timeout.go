package uti

import "time"

func TimeOut(duration time.Duration) {
	go func() {
		<-time.After(duration)
		panic("timeout")
	}()
}
