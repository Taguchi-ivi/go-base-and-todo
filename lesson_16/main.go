package main

import (
	"context"
	"fmt"
	"time"
)

// flag
// func main() {
// 	// コマンドラインのオプション処理
// 	// 存在しないオプションを指定するとhelpメッセージが表示される
// 	// go run main.go -n 20 -m message -x
// 	var (
// 		max int
// 		msg string
// 		x   bool
// 	)

// 	flag.IntVar(&max, "n", 32, "処理数の最大値")
// 	flag.StringVar(&msg, "m", "", "処理メッセージ")
// 	flag.BoolVar(&x, "x", false, "拡張オプション")

// 	flag.Parse()
// 	fmt.Println(max, msg, x)
// }

// logger
// func main() {
// 	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
// 	logger.Println("message")
// 	logger.Printf("%#v\n", logger)
// }

// sync
// go routineの同期処理
// func main() {
// 	// sync.WaitGroupを生成
// 	wg := new(sync.WaitGroup)
// 	// 待ち受けるgoroutineの数を追加(3つ)
// 	wg.Add(3)

// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			fmt.Println("1st goroutine")
// 		}
// 		wg.Done() // 完了
// 	}()
// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			fmt.Println("2st goroutine")
// 		}
// 		wg.Done()
// 	}()
// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			fmt.Println("3st goroutine")
// 		}
// 		wg.Done()
// 	}()

// 	// すべてのgoroutineが終了するまで待機
// 	// Doneが3つ呼ばれるまで待機
// 	wg.Wait()
// }

// context
// 値の受け渡し(user idやrequest idなど), キャンセル処理, タイムアウト処理など
func longProcess(ctx context.Context, ch chan string) {
	fmt.Println("start ")
	time.Sleep(2 * time.Second)
	fmt.Println("end")
	ch <- "実行結果"
}

func main() {
	ch := make(chan string)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	go longProcess(ctx, ch)

L:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("error")
			fmt.Println(ctx.Err())
			break L
		case s := <-ch:
			fmt.Println("success")
			fmt.Println(s)
			break L
		}
	}
	fmt.Println("main関数終了")
}
