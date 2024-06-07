package main

import (
	"fmt"
	"time"
)

func DoWorkVer1(done <-chan interface{}, palseInterval time.Duration) (<-chan interface{}, <-chan time.Time) {
	heartbeat := make(chan interface{}, 1)
	results := make(chan time.Time)

	go func() {
		defer close(heartbeat)
		defer close(results)

		heartbeatPulse := time.Tick(palseInterval)
		workGen := time.Tick(2 * time.Second)

		sendHeartbeatPalse := func() {
			select {
			case heartbeat <- struct{}{}:
			default:
			}
		}

		sendResult := func(result time.Time) {
			for {
				select {
				case <-done:
					return
				case <-heartbeatPulse:
					sendHeartbeatPalse()
				case results <- result.Local():
					return
				}
			}
		}

		for {
			select {
			case <-done:
				return
			case <-heartbeatPulse:
				sendHeartbeatPalse()
			case result := <-workGen:
				sendResult(result)
			}
		}
	}()

	return heartbeat, results
}

func main() {
	// heartbeat01
	// 並行処理の動作が正常に行われていることを外部に通知するための方法
	done := make(chan interface{})
	const timeout = 2 * time.Second

	heartbeat, result := DoWorkVer1(done, timeout/2)
	for {
		select {
		case _, ok := <-heartbeat:
			if !ok {
				return
			}
			fmt.Println("received heartbeat")
		case r, ok := <-result:
			if !ok {
				return
			}
			fmt.Printf("received result: %v\n", r)
		case <-time.After(timeout):
			fmt.Println("worker goroutine is dead")
			return
		}
	}
}
