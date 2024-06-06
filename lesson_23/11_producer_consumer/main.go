package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func worker(workerNum int, ch <-chan int, wg *sync.WaitGroup) {
	for num := range ch {
		time.Sleep(1 * time.Second)
		fmt.Printf("worker%d: %d\n", workerNum, num)
		wg.Done()
	}
}

func main() {
	// producer consumer
	// タスクの生成と処理を分離する
	ch := make(chan int)
	var wg sync.WaitGroup

	for i := 1; i <= runtime.NumCPU(); i++ {
		go worker(i, ch, &wg)
	}

	for i := 1; i < 100; i++ {
		wg.Add(1)
		go func(x int) {
			// producerとしてchに値を送信
			ch <- x * 2
		}(i)
	}
	wg.Wait()
	close(ch)
}
