package main

import "fmt"

// pipeline　データを受け取って、何らかの処理を行ってから、またどこかにデータを送る
// 何らかの処理 => 何らかの処理 => 何らかの処理
// データを流すことをステージと呼ぶ

// pipelineのステージを関数として定義する
// func double(sl []int) []int {
// 	doubleSlice := make([]int, 0, len(sl))
// 	for i, v := range sl {
// 		doubleSlice[i] = v * 2
// 	}
// 	return doubleSlice
// }

// pipelineのステージを関数として定義する
// func add(sl []int) []int {
// 	addSlice := make([]int, 0, len(sl))
// 	for i, v := range sl {
// 		addSlice[i] = v + 1
// 	}
// 	return addSlice
// }

func generator(done <-chan interface{}, integers ...int) <-chan int {
	intStream := make(chan int)

	go func() {
		defer close(intStream)
		for _, v := range integers {
			select {
			case <-done:
				return
			case intStream <- v:
			}
		}
	}()
	return intStream
}

// 並行処理、goroutineとchannelを使って定義
func double(done <-chan interface{}, intStream <-chan int) <-chan int {
	doubleChan := make(chan int)
	go func() {
		defer close(doubleChan)
		for v := range intStream {
			select {
			case <-done:
				return
			case doubleChan <- v * 2:
			}
		}
	}()
	return doubleChan
}

// 並行処理、goroutineとchannelを使って定義
func add(done <-chan interface{}, intStream <-chan int) <-chan int {
	addChan := make(chan int)
	go func() {
		defer close(addChan)
		for v := range intStream {
			select {
			case <-done:
				return
			case addChan <- v + 1:
			}
		}
	}()
	return addChan
}

func main() {
	// ints := []int{1, 2, 3, 4, 5}

	// addステージに渡す -> doubleステージに渡す
	// これにより独立することと、doubleとaddを使いやすくなる
	// 並行処理をしたい場合は、やはりgoroutineとchannelを使とよい
	// for _, v := range double(add(ints)) {
	// 	// for _, v := range double(double(add(ints))) {
	// 	fmt.Println("v: ", v)
	// }

	// 並行処理、goroutineとchannelを使って定義
	done := make(chan interface{})
	defer close(done)
	intStream := generator(done, 1, 2, 3, 4, 5)
	for v := range double(done, add(done, intStream)) {
		// for _, v := range double(double(add(ints))) {
		fmt.Println("v: ", v)
	}
}
