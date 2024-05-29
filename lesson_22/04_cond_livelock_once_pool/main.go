package main

func main() {
	// cond
	// var mutex sync.Mutex
	// cond := sync.NewCond(&mutex)

	// for _, name := range []string{"A", "B", "C"} {
	// 	go func(name string) {
	// 		mutex.Lock()
	// 		defer mutex.Unlock()

	// 		// ここで待機(シグナルが来るまで待機)
	// 		cond.Wait()
	// 		fmt.Println(name)
	// 	}(name)
	// }

	// fmt.Println("Ready...")
	// time.Sleep(time.Second)
	// fmt.Println("Go!")
	// // 一個ずつシグナルを送る
	// // for i := 0; i < 3; i++ {
	// // 	time.Sleep(time.Second)
	// // 	cond.Signal()
	// // }
	// cond.Broadcast() // すべてのgoroutineにシグナルを送る
	// time.Sleep(time.Second)
	// fmt.Println("Done!")

	// livelock
	// condにて、シグナルを送る前に、待機しているgoroutineがいない場合、deadlockのような状態になること
	// こちらも並行処理の注意点 (処理は省く)

	// once
	// Doで渡された関数は一度だけ実行する
	// 関数名が違っても、同じ関数として扱われる
	// count := 0
	// increment := func() {
	// 	count++
	// }
	// decrement := func() {
	// 	count--
	// }
	// var once sync.Once
	// once.Do(increment)
	// once.Do(decrement)
	// fmt.Println(count)

	// pool
	// オブジェクトのキャッシュを管理する
	// 定義したmypoolの中に、存在しないインスタンスを取得しようとすると、Newで定義した関数を実行してインスタンスを生成する
	// mypool := &sync.Pool{
	// 	New: func() interface{} {
	// 		fmt.Println("Create new instance.")
	// 		return struct{}{}
	// 	},
	// }
	// instance := mypool.Get()
	// mypool.Put(instance)
	// // ここではキャッシュされたインスタンスが返される
	// mypool.Get()
	// type Person struct {
	// 	Name string
	// }
	// instance3の時に初めてNewが実行される
	// mypool.Put(Person{Name: "1"})
	// mypool.Put(Person{Name: "2"})
	// instance1 := mypool.Get()
	// instance2 := mypool.Get()
	// instance3 := mypool.Get()
	// fmt.Println(instance1, instance2, instance3)

	// connection poolテストを main_test.go に記述
}
