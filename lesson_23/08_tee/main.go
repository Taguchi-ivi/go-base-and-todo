package main

import "fmt"

func generator(done <-chan interface{}) <-chan interface{} {
	intChan := make(chan interface{})
	go func() {
		defer close(intChan)
		defer fmt.Println("generator done")
		for i := 0; i < 10; i++ {
			select {
			case <-done:
				return
			case intChan <- i:
			}
		}
	}()
	return intChan
}

func orDone(done, c <-chan interface{}) <-chan interface{} {
	valChan := make(chan interface{})
	go func() {
		defer close(valChan)
		for {
			select {
			case <-done:
				return
			case val, ok := <-c:
				if !ok {
					return
				}
				select {
				case valChan <- val:
				case <-done:
				}
			}
		}
	}()
	return valChan
}

func tee(done, c <-chan interface{}) (<-chan interface{}, <-chan interface{}) {
	out1 := make(chan interface{})
	out2 := make(chan interface{})
	go func() {
		defer close(out1)
		defer close(out2)
		for val := range orDone(done, c) {
			var Out1, Out2 = out1, out2
			for i := 0; i < 2; i++ {
				select {
				// 書き込んだらnilにすることで、もう一方のチャネルには書き込まないようにする
				case Out1 <- val:
					Out1 = nil
				case Out2 <- val:
					Out2 = nil
				}
			}
		}
	}()
	return out1, out2
}

func main() {
	// tee: チャネルからの値を二つに分ける
	// チャネルから受信した値を二つの分岐して扱いたいときに便利な機能
	done := make(chan interface{})

	out1, out2 := tee(done, generator(done))
	for v1 := range out1 {
		fmt.Printf("out1: %v, out2: %v\n", v1, <-out2)
	}
}
