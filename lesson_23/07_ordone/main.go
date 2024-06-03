package main

import (
	"fmt"
	"time"
)

func generator() <-chan int {
	intStream := make(chan int)

	go func() {
		defer close(intStream)

		for i := 0; i < 10; i++ {
			intStream <- i
		}
	}()

	return intStream
}

// done channel
func signal(after time.Duration) <-chan interface{} {
	done := make(chan interface{})

	go func() {
		defer close(done)
		defer fmt.Println("signal close")

		time.Sleep(after)
	}()

	return done
}

func orDone(done <-chan interface{}, c <-chan int) <-chan interface{} {
	valStream := make(chan interface{})

	go func() {
		defer close(valStream)
		defer fmt.Println("orDone close")

		for {
			select {
			case <-done:
				return
			case v, ok := <-c:
				if !ok {
					return
				}
				select {
				case valStream <- v:
				case <-done:
				}
			}
		}
	}()

	return valStream

}

func main() {
	// orDone どちらかが終了したら閉じる
	// 終了判定のチャネルと入力データのチャネルどちらかが閉じたら終了する
	start := time.Now()
	done := signal(10 * time.Second)
	intStream := generator()

	// 下記だと 使いたい部分のコード量が多くなってしまう
	// loop:
	// 	for {
	// 		select {
	// 		case <-done:
	// 			fmt.Println("done")
	// 			break loop
	// 		case v, ok := <-intStream:
	// 			if !ok {
	// 				fmt.Println("intStream close")
	// 				return
	// 			}
	// 			fmt.Println(v)
	// 		}
	// 	}

	// or done
	for val := range orDone(done, intStream) {
		fmt.Println(val)
	}

	fmt.Println(time.Since(start))
}
