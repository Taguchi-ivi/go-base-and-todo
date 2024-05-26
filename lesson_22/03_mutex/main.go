package main

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	mu    sync.Mutex
	value int
	name  string
}

func main() {
	// mutex 競合状態
	// var wg sync.WaitGroup
	// var memoryAccess sync.Mutex
	// var data int

	// wg.Add(1)

	// go func() {
	// 	defer wg.Done()
	// 	// この処理が終わるまで他のgoroutineが実行されない
	// 	// lockが解除されるまで他のgoroutineが実行されない
	// 	// lock, unlockはコストがかかるので、必要なときだけ使う
	// 	memoryAccess.Lock()
	// 	data++
	// 	memoryAccess.Unlock()
	// }()
	// wg.Wait()

	// memoryAccess.Lock()
	// if data == 0 {
	// 	fmt.Println(0)
	// } else {
	// 	fmt.Println(data)
	// }
	// memoryAccess.Unlock()

	// deadlock
	// var wg sync.WaitGroup
	// printSum := func(v1, v2 *value) {
	// 	defer wg.Done()
	// 	v1.mu.Lock()
	// 	fmt.Printf("%v is locked\n", v1.name)
	// 	defer v1.mu.Unlock()

	// 	time.Sleep(2 * time.Second)

	// 	v2.mu.Lock()
	// 	fmt.Printf("%v is locked\n", v2.name)
	// 	defer v2.mu.Unlock()

	// 	println(v1.value + v2.value)
	// }

	// なぜdeadlockになるか
	// 1. printSum(a, b)が実行される
	// 2. a.mu.Lock()が実行される
	// 3. time.Sleep(2 * time.Second)が実行される
	// 4. printSum(b, a)が実行される
	// 5. b.mu.Lock()が実行される
	// 6. a.mu.Lock()が解除されるまで、b.mu.Lock()が実行されない
	// 7. a.mu.Lock()が解除されるまで、printSum(b, a)が実行されない
	// 8. a.mu.Lock()が解除されるまで、printSum(a, b)が実行されない
	// 9. a.mu.Lock()が解除されるまで、printSum(b, a)が実行されない

	// var a, b value
	// a.name = "a"
	// b.name = "b"
	// wg.Add(2)
	// go printSum(&a, &b)
	// go printSum(&b, &a)
	// wg.Wait()

	// リソースの枯渇(あるプロセスが他のプロセスが使いたいリソースを使っている)
	// 正しくリソースを使うこと。リソースを使い終わったら解放すること。
	// 正解はない。処理速度とlockのコストを考慮して使うことが重要になる
	// var wg sync.WaitGroup
	// var lock sync.Mutex

	// const timer = 1 * time.Second
	// greedWorker := func() {
	// 	defer wg.Done()
	// 	count := 0
	// 	begin := time.Now()
	// 	for time.Since(begin) <= timer {
	// 		lock.Lock()
	// 		time.Sleep(3 * time.Nanosecond)
	// 		lock.Unlock()
	// 		count++
	// 	}
	// 	fmt.Println("g count:", count)
	// }

	// politeWorker := func() {
	// 	defer wg.Done()
	// 	count := 0
	// 	begin := time.Now()
	// 	for time.Since(begin) <= timer {
	// 		lock.Lock()
	// 		time.Sleep(1 * time.Nanosecond)
	// 		lock.Unlock()
	// 		lock.Lock()
	// 		time.Sleep(1 * time.Nanosecond)
	// 		lock.Unlock()
	// 		lock.Lock()
	// 		time.Sleep(1 * time.Nanosecond)
	// 		lock.Unlock()
	// 		count++
	// 	}
	// 	fmt.Println("p count:", count)
	// }

	// wg.Add(2)
	// go greedWorker()
	// go politeWorker()
	// wg.Wait()

	// mutex,RWMutex
	// mutexは読み込みも書き込みもlockがかかる
	// RWMutexは読み込みの制御と書き込みの制御を分けることができる
	// mutexを使うと、読み込みも書き込みもlockがかかるので、最大5*2=10秒かかる
	// RWMutexを使い、一部の読み込みをlockしないようにすると、最大5+1(同時に行われる)=6秒で終わる
	var count int
	// var lock sync.Mutex
	var lock sync.RWMutex
	var wg sync.WaitGroup

	increment := func(wg *sync.WaitGroup, l sync.Locker) {
		l.Lock()
		defer l.Unlock()
		defer wg.Done()

		fmt.Println("increment")
		count++
		time.Sleep(1 * time.Second)
	}

	read := func(wg *sync.WaitGroup, l sync.Locker) {
		l.Lock()
		defer l.Unlock()
		defer wg.Done()

		fmt.Println("read")
		time.Sleep(1 * time.Second)
	}

	start := time.Now()
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go increment(&wg, &lock)
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		// go read(&wg, &lock)
		go read(&wg, lock.RLocker())
	}
	wg.Wait()
	end := time.Now()
	fmt.Println("end.sub(start):", end.Sub(start))
}
