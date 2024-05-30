package main

import (
	"fmt"
	"log"
	"os"
)

type Result struct {
	Response *os.File
	Error    error
}

// func CheckFiles(done <-chan interface{}, files ...string) <-chan *os.File {
func CheckFiles(done <-chan interface{}, files ...string) <-chan Result {
	// res := make(chan *os.File)
	res := make(chan Result)

	go func() {
		defer close(res)
		for _, file := range files {
			var r Result
			file, err := os.Open(file)
			// if err != nil {
			// 	fmt.Println(err)
			// 	return
			// }
			r = Result{Response: file, Error: err}

			select {
			case <-done:
				return
				// case res <- file:
			case res <- r:
			}
		}
	}()

	return res
}

func main() {
	// 並行処理でのエラーハンドリング エラー処理を、どのgoroutineで行うか切り分ける
	// 誰が、どの関数がエラーの責任を負うかを考える必要がある。
	// この例では、このmain関数がエラーの責任を負うべき.関数に任せないこと
	done := make(chan interface{})

	defer close(done)

	file := []string{"main.go", "x.go", "y.go"}
	for res := range CheckFiles(done, file...) {
		// fmt.Println(res)
		if res.Error != nil {
			log.Println("Error:", res.Error)
			continue
		}

		fmt.Println("response:", res.Response.Name())
	}
}
