package main

import (
	"fmt"
	"math/rand"
)

func DoWorkVer2(done <-chan interface{}) (<-chan interface{}, <-chan int) {
	heartbeat := make(chan interface{}, 1)
	workStream := make(chan int)

	// 仕事ごと、goroutineごとにハートビートを送信する
	go func() {
		defer close(heartbeat)
		defer close(workStream)
		for i := 0; i < 10; i++ {
			select {
			case heartbeat <- struct{}{}: // 空の構造体ハートビートを送信
			default:
			}
			select {
			case <-done:
				return
			case workStream <- rand.Intn(10):
			}
		}
	}()
	return heartbeat, workStream
}

func main() {
	// heartbeat02
	// 仕事の単位ごとにハートビート(鼓動)を送信する
	// 鼓動 = 並行処理の動作が正常に行われていることを外部に通知するための方法
	done := make(chan interface{})
	defer close(done)

	heartbeat, results := DoWorkVer2(done)
	// 仕事ごとにハートビートが送信される (pulseが表示される)
	for {
		select {
		case _, ok := <-heartbeat:
			if ok {
				fmt.Println("pulse")
			} else {
				return
			}
		case result, ok := <-results:
			if ok {
				fmt.Printf("result %v\n", result)
			} else {
				return
			}
		}
	}
}
