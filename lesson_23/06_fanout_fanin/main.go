package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func repeatFn(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		for {
			select {
			case <-done:
				return
			case valueStream <- fn():
			}
		}
	}()
	return valueStream
}

func take(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
	takeStream := make(chan interface{})
	go func() {
		defer close(takeStream)
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case takeStream <- <-valueStream:
			}
		}
	}()
	return takeStream
}

func toInt(done <-chan interface{}, valueStream <-chan interface{}) <-chan int {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for v := range valueStream {
			select {
			case <-done:
				return
			case intStream <- v.(int):
			}
		}
	}()
	return intStream
}

// 素数を見つける
func primeFinder(done <-chan interface{}, intStream <-chan int) <-chan interface{} {
	primeStream := make(chan interface{})
	go func() {
		defer close(primeStream)
	L:
		for i := range intStream {
			for div := 2; div <= i/2; div++ {
				if i%div == 0 {
					continue L
				}
			}
			select {
			case <-done:
				return
			case primeStream <- i:
			}
		}
	}()
	return primeStream
}

func random() interface{} {
	return rand.Intn(50000000)
}

func fanIn(done <-chan interface{}, channels ...<-chan interface{}) <-chan interface{} {
	var wg sync.WaitGroup
	multiplexedStream := make(chan interface{})
	multiplex := func(c <-chan interface{}) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case multiplexedStream <- i:
			}
		}
	}
	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c)
	}
	go func() {
		wg.Wait()
		close(multiplexedStream)
	}()
	return multiplexedStream
}

func main() {
	// ファンアウト、ファンイン
	// ファンアウト: 一つのチャネルから複数のチャネルにデータを送信する
	// ファンイン: 複数のチャネルからのデータを一つのチャネルにまとめる
	// 1本のpipelineの流れを複数に分ける処理(同時に複数のgoroutineを起動する処理)
	// 複数のgoroutineで処理されたchannelを1本に集約する処理をファンインと呼ぶ

	done := make(chan interface{})
	defer close(done)

	randIntStream := toInt(done, repeatFn(done, random))

	start := time.Now()

	// ファンアウト,ファンイン pipelineだと18秒かかる処理が5秒とかになる
	numFinders := runtime.NumCPU()
	fmt.Println("prime finders :", numFinders)
	finders := make([]<-chan interface{}, numFinders)
	for i := 0; i < numFinders; i++ {
		finders[i] = primeFinder(done, randIntStream)
	}
	for prime := range take(done, fanIn(done, finders...), 10) {
		fmt.Println(prime)
	}

	// pipelineのステージを関数として定義する
	// for prime := range take(done, primeFinder(done, randIntStream), 10) {
	// 	fmt.Println(prime)
	// }

	fmt.Println(time.Since(start))
}
