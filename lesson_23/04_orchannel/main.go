package main

import (
	"fmt"
	"time"
)

// or channel
// 複数のdoneチャネルからのデータを受け取り、一つのチャネルにまとめる
// どちらかのチャネルがデータを送信すると、そのデータを受け取る(複数のうちどれか一つでもデータを受け取ると終了する)

func or(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	orDone := make(chan interface{})
	go func() {
		defer close(orDone)

		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
				// 3つ以上の場合は、残りのチャネルをスライスでまとめて再起的に呼び出す
			case <-or(append(channels[3:], orDone)...):
			}
		}
	}()
	return orDone
}

func signal(after time.Duration) <-chan interface{} {
	done := make(chan interface{})
	go func() {
		defer close(done)
		time.Sleep(after)
	}()
	return done
}

func main() {
	start := time.Now()
	// どれか一つでもデータを受け取ると終了する
	<-or(signal(time.Hour), signal(time.Minute), signal(time.Second), signal(time.Millisecond))
	fmt.Printf("done after %v\n", time.Since(start))
}
