package main

import (
	"context"
	"fmt"
	"time"
)

// withTimeout
// func shortProcess(done <-chan interface{}) (bool, error) {
// func shortProcess(ctx context.Context) (bool, error) {
// 	shortWork := time.NewTicker(1 * time.Second)

// 	ctx, cancel := context.WithTimeout(ctx, 1*time.Millisecond)
// 	defer cancel()

// 	select {
// 	// case <-done:
// 	case <-ctx.Done():
// 		return false, fmt.Errorf("short process canceled\n")
// 	case <-shortWork.C:
// 	}
// 	return true, nil
// }

// deadline
func shortProcess(ctx context.Context) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	switch isDone, err := doSomething(ctx); {
	case err != nil:
		return false, err
	case isDone:
		return isDone, nil
	}
	return false, fmt.Errorf("unSupported error\n")
}

// func longProcess(done <-chan interface{}) (bool, error) {
func longProcess(ctx context.Context) (bool, error) {
	shortWork := time.NewTicker(5 * time.Second)

	select {
	// case <-done:
	case <-ctx.Done():
		return false, ctx.Err()
	case <-shortWork.C:
	}
	return true, nil
}

func doSomething(ctx context.Context) (bool, error) {
	if deadline, ok := ctx.Deadline(); ok {
		if deadline.Sub(time.Now().Add(2*time.Second)) <= 0 {
			return false, context.DeadlineExceeded
		}
	}
	select {
	case <-ctx.Done():
		return false, ctx.Err()
	case <-time.After(3 * time.Second):
	}
	return true, nil
}

type ctxKey int

const (
	ctxUserID ctxKey = iota
	ctxAuthToken
)

func Set(userId, authToken string) context.Context {
	ctx := context.Background()
	// ctx = context.WithValue(ctx, "userID", userId)
	// ctx = context.WithValue(ctx, "authToken", authToken)
	ctx = context.WithValue(ctx, ctxUserID, userId)
	ctx = context.WithValue(ctx, ctxAuthToken, authToken)
	return ctx
}

func Get(ctx context.Context) (string, string) {
	userId := ctx.Value(ctxUserID).(string)
	authToken := ctx.Value(ctxAuthToken).(string)
	return userId, authToken
}

func main() {
	// context cancel
	// そもそもcontextとは
	// Deadline, Cancel, Timeout, Value
	// デッドラインが設定されている時はその時刻を返す。設定されていない場合はokがfalseになる
	// Done() contextがキャンセルまたはタイムアウトじにcloseされる
	// err goroutineがキャンセルされた時にキャンセルの理由を返す
	// value ランタイムの割り込みに関数情報に加えてリクエストに応じた情報が渡される必要がある場合に使われる keyに紐づいた値を取得する
	// からのcontextを作成する background(), TODO()はほぼ本番では使わないらしい
	// var wg sync.WaitGroup
	// ctx, cancel := context.WithCancel(context.Background())

	// defer cancel()

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	if isDone, err := shortProcess(ctx); err != nil {
	// 		fmt.Printf("short process error: %v\n", err)
	// 		fmt.Println(isDone)
	// 		cancel()
	// 		return
	// 	}
	// 	fmt.Println("short process done")
	// }()

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	if isDone, err := longProcess(ctx); err != nil {
	// 		fmt.Printf("long process error: %v\n", err)
	// 		fmt.Println(isDone)
	// 		cancel()
	// 		return
	// 	}
	// 	fmt.Println("long process done")
	// }()

	// value
	ctx := Set("user1", "authToken1")
	userId, authToken := Get(ctx)
	fmt.Printf("userId: %s, authToken: %s\n", userId, authToken)
}
