package main

// var data []int = []int{1, 2, 3, 4, 5}

// dataはどこからでもアクセスできるが、writeToChan関数内でのみアクセスするルールを作ることで、アクセス権を拘束する
// func writeToChan(writeChan chan<- int) {
// 	defer close(writeChan)
// 	for _, v := range data {
// 		writeChan <- v
// 	}
// }

func chanOwner() <-chan int {
	resultCh := make(chan int, 5)
	go func() {
		defer close(resultCh)
		for i := 1; i <= 5; i++ {
			resultCh <- i
		}
	}()

	return resultCh
}

func consumer(resultCh <-chan int) {
	for v := range resultCh {
		println(v)
	}
	println("Done!")
}

func main() {
	// データのアクセス権を拘束するパターン アクセス範囲を縛る
	// ルールで対応するパターン
	// handleData := make(chan int)
	// go writeToChan(handleData)
	// for v := range handleData {
	// 	fmt.Println(v)
	// }

	// 複数の並行プロセスに公開するパターン
	// このパターンは、chanOwner関数内で生成したチャネルをconsumer関数に渡すことで、chanOwner関数内で生成したチャネルのアクセス権をconsumer関数に移譲している
	// 読み込みと書き込みの権限を分離することで、データの競合を防ぐ
	// 同期する際のコストを最小限に抑えることができる
	results := chanOwner()
	consumer(results)
}
