package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readGoFile(path string) chan string {
	promise := make(chan string)

	go func() {
		defer close(promise)
		content, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Printf("read error: %v\n", err)
		} else {
			promise <- string(content)
		}
	}()
	return promise
}

func printFunc(futureSource chan string) chan []string {
	promise := make(chan []string)
	go func() {
		defer close(promise)
		var result []string

		for _, line := range strings.Split(<-futureSource, "\n") {
			if strings.HasPrefix(line, "func ") {
				result = append(result, line)
			}
		}
		promise <- result
	}()
	return promise
}

func main() {
	// future promise
	// タスク分割の手法の一つ
	futureSource := readGoFile("main.go")
	futureFuncs := printFunc(futureSource)

	fmt.Println(strings.Join(<-futureFuncs, "\n"))
}
