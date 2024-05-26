package main

import (
	"fmt"
	"sync"
)

func sayHello(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello")
}

// 並行処理 基礎Goroutine
func main() {

	// Goroutine基礎
	// var wg sync.WaitGroup
	// wg.Add(1)

	// go sayHello(&wg)

	// wg.Add(1)
	// // 無名関数
	// go func() {
	// 	defer wg.Done()
	// 	fmt.Println("こんにちは")
	// }()
	// wg.Wait()

	// クロージャとforloopを使った注意点
	var wg sync.WaitGroup
	say := "hello"
	tasks := []string{"a", "b", "c"}
	wg.Add(1)
	go func() {
		defer wg.Done()
		// 上書きされる
		say = "こんにちは"
	}()
	for _, task := range tasks {
		wg.Add(1)
		// 依存する変数を引数に渡さないと、ループが終わった時に最後の値が使われる
		go func(task string) {
			defer wg.Done()
			fmt.Println(task)
		}(task)
	}
	wg.Wait()
	fmt.Println("say:", say)

}
