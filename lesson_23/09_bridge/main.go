package main

import "fmt"

func generateVals() <-chan <-chan interface{} {
	chanStream := make(chan (<-chan interface{}))
	go func() {
		defer close(chanStream)

		// チャネルのチャネルを生成 チャネルの中にチャネルを入れる
		for i := 0; i < 10; i++ {
			stream := make(chan interface{}, 1)
			stream <- i
			close(stream)
			chanStream <- stream
		}
	}()
	return chanStream
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

func bridge(done <-chan interface{}, chanStream <-chan <-chan interface{}) <-chan interface{} {
	valChan := make(chan interface{})
	go func() {
		defer close(valChan)
		for {
			// チャネルのチャネルからチャネルを取り出す
			var stream <-chan interface{}
			select {
			case maybeStream, ok := <-chanStream:
				if !ok {
					return
				}
				stream = maybeStream
			case <-done:
				return
			}

			for val := range orDone(done, stream) {
				select {
				case valChan <- val:
				case <-done:
				}
			}
		}
	}()
	return valChan
}

func main() {
	// Bridgeチャンネル channelのchannel
	// channelのchannelを単一のchannelとして扱う
	done := make(chan interface{})

	for v := range bridge(done, generateVals()) {
		fmt.Println(v)
	}
}
