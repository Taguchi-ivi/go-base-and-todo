package main

import (
	"fmt"
	"sync"
	"time"
)

// 並行処理 sync,wait
func main() {
	var wg sync.WaitGroup

	// countを1つ増やす場合は必ず監視対象の外側でAddを実行する main処理がendされる前に必ずaddしないといけないから
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("1st Goroutine start")
		time.Sleep(1 * time.Second)
		fmt.Println("1st Goroutine end")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2st Goroutine start")
		time.Sleep(1 * time.Second)
		fmt.Println("2st Goroutine end")
	}()

	wg.Wait()
}
