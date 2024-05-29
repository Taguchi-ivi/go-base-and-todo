package main

func main() {
	// // select
	// // 複数のチャネルの操作をまとめるもの
	// a := make(chan int)
	// b := make(chan int)
	// close(a)
	// // bが閉じられていないのでaが待ってしまい、デッドロックになる
	// // <-b
	// // <-a
	// // selectを使うことで、どちらかのチャネルが受信できるまで待つ状態を作れる
	// select {
	// case <-b:
	// case <-a:
	// }

	// 	ch1 := make(chan int)
	// 	ch2 := make(chan int)
	// 	done := make(chan interface{})
	// 	go func() {
	// 		time.Sleep(2 * time.Second)
	// 		close(done)
	// 	}()
	// 	go func() {
	// 		defer close(ch1)
	// 		for i := 0; i < 10; i++ {
	// 			time.Sleep(1 * time.Second)
	// 			ch1 <- i
	// 		}
	// 	}()
	// 	go func() {
	// 		defer close(ch2)
	// 		for i := 0; i < 10; i++ {
	// 			time.Sleep(1 * time.Second)
	// 			ch2 <- i
	// 		}
	// 	}()
	// loop:
	// 	for {
	// 		select {
	// 		case v, ok := <-ch1:
	// 			if !ok {
	// 				break loop
	// 			}
	// 			println("ch1:", v)
	// 		case v, ok := <-ch2:
	// 			if !ok {
	// 				break loop
	// 			}
	// 			println("ch2:", v)
	// 		case <-done:
	// 			break loop
	// 		case <-time.After(3 * time.Second):
	// 			println("timeout")
	// 			break loop
	// 		}
	// 	}

	// nilチャネルは常にブロックする デッドロックになる
	var ch <-chan int
	select {
	case <-ch:
	default: // デフォルトの処理を入れることでデッドロックを回避できる
		println("default")
	}
}
