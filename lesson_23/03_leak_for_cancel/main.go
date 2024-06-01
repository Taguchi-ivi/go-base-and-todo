package main

import (
	"fmt"
	"math/rand"
)

func DoSomethingForRead(done <-chan interface{}, strings <-chan string) <-chan interface{} {
	completed := make(chan interface{})

	go func() {
		defer fmt.Println("doSomething done")
		defer close(completed)

		for {
			select {
			case s := <-strings:
				fmt.Println(s)
			case <-done:
				return
			}
		}
	}()

	return completed
}

// func DoSomethingForWrite() <-chan int {
func DoSomethingForWrite(done chan interface{}) <-chan int {
	readStream := make(chan int)

	go func() {
		defer fmt.Println("doSomething done")
		defer close(readStream)

		for {
			readStream <- rand.Intn(100)
			select {
			case readStream <- rand.Intn(100):
			case <-done:
				return
			}
		}
	}()

	return readStream
}

// goroutineリークをキャンセル処理で対策
// goroutineはキャンセル処理を行わないと、メモリが解放されないため、goroutineリークが発生する可能性がある。
func main() {
	// goroutineリークをキャンセル処理で対策 Read(読み込み) ver
	// done := make(chan interface{})
	// // nilチャネルを渡すと、goroutineリークが発生する
	// // DoSomething(nil)
	// complated := DoSomethingForRead(done, nil)
	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	close(done)
	// }()
	// <-complated
	// fmt.Println("main done")

	// goroutineリークをキャンセル処理で対策 Write(書き込み) ver
	// プロセスが生き残る限り、メモリは解放されない
	// キャンセルのシグナルを送ることで、メモリを解放する
	done := make(chan interface{})
	// readStream := DoSomethingForWrite()
	readStream := DoSomethingForWrite(done)
	for i := 1; i <= 3; i++ {
		fmt.Println(<-readStream)
	}
	close(done)
	fmt.Println("main done")

	// memo
	// goroutineを生成したら、必ずキャンセル処理も行うという意識を持つ

}
