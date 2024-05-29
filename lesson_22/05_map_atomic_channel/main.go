package main

func main() {
	// map error
	// mapは参照型であり、複数のgoroutineから同時にアクセスされる場合、データ競合が発生する可能性がある
	// concurrent map writesのエラーが発生する可能性がある
	// var wg sync.WaitGroup
	// m := map[string]int{"a": 1, "b": 2}

	// for i := 0; i < 10; i++ {
	// 	wg.Add(2)
	// 	go func() {
	// 		defer wg.Done()
	// 		m["a"] = rand.Intn(100)
	// 		m["b"] = rand.Intn(100)
	// 	}()
	// 	go func() {
	// 		defer wg.Done()
	// 		m["a"] = rand.Intn(100)
	// 		m["b"] = rand.Intn(100)
	// 	}()
	// }
	// wg.Wait()
	// sync.map 複数のgoroutineから安全にアクセスできる
	// smap := &sync.Map{}
	// smap.Store("hello", "world") // key: hello, value: world
	// smap.Store(1, 2)             // key: 1, value: 2
	// smap.Range(func(key, value interface{}) bool {
	// 	println(key, value)
	// 	return true
	// })
	// smap.Delete(1)
	// smap.Range(func(key, value interface{}) bool {
	// 	println(key, value)
	// 	return true
	// })
	// v, ok := smap.Load("hello")
	// if ok {
	// 	println(v)
	// }
	// smap.LoadOrStore("hello", "woooorld") // keyが存在しない場合、valueを追加する

	// atomic
	// atomicを使うことによって、データ競合を防ぐことができる(mutexよりも高速)
	// var count int64
	// increment := func() {
	// 	atomic.AddInt64(&count, 1)
	// }
	// increment()
	// fmt.Println("Count:", count)

	// channel 同期処理を行う FIFO(First In First Out)のデータ構造
	// channelには3つの種類がある(バッファ付き、バッファなし、双方向)
	// ch := make(chan string)
	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	ch <- "Hello!"
	// }()
	// fmt.Println("Received:", <-ch)

	// channel close
	// ch := make(chan string)
	// go func() {
	// 	ch <- "Hello!"
	// }()
	// v, ok := <-ch
	// fmt.Println("Received:", v, ok)
	// close(ch)
	// // close後は、channelに値がない場合、第2戻り値がfalseになる
	// v, ok = <-ch
	// fmt.Println("Received:", v, ok)
	// ch := make(chan int)
	// go func() {
	// 	defer close(ch) // closeしないと、rangeでエラーが発生する(deadlock) 値を待ち続けてしまうので
	// 	for i := 0; i < 10; i++ {
	// 		ch <- i
	// 	}
	// }()
	// for v := range ch {
	// 	fmt.Println("Received:", v)
	// }

	// channel buffer
	// バッファを付けないと、読み込みがない場合、書き込みがブロックされる(書き込みと読み込みが交互に行われる)
	// バッファを付けることで、書き込みがブロックされることなく、書き込みを先に対応して、その後読み込むことができる
	// ch := make(chan int)
	// ch := make(chan int, 5)
	// go func() {
	// 	defer fmt.Println("Producer done.")
	// 	defer close(ch)
	// 	for i := 0; i < 5; i++ {
	// 		fmt.Println("writing:", i)
	// 		ch <- i
	// 	}
	// }()
	// for v := range ch {
	// 	time.Sleep(1 * time.Second)
	// 	fmt.Println("Read:", v)
	// }

	// channelのライフサイクルのカプセル化
	// channelの所有権を持つ関数を作成し、その関数内でchannelを生成し、返すことで、channelのライフサイクルをカプセル化する
	// 必ず初期化と終了処理を行うことで、channelのリークを防ぐ
	// chanOwner := func() <-chan int {
	// 	resultCh := make(chan int, 5)
	// 	go func() {
	// 		defer close(resultCh)
	// 		for i := 0; i < 5; i++ {
	// 			resultCh <- i
	// 		}
	// 	}()
	// 	return resultCh
	// }
	// resultCh := chanOwner()
	// for result := range resultCh {
	// 	fmt.Println("Received:", result)
	// }
	// fmt.Println("Done.")

}
