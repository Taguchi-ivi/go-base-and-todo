package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"sync"
	"testing"
	"time"
)

// connection poolを用いることで、パフォーマンスを向上させることができる
// 生成したインスタンスをキャッシュしておき、再利用することで、生成コストを削減する

func connetcToService() interface{} {
	time.Sleep(1 * time.Second)
	return struct{}{}
}

func warmServiceConnCache() *sync.Pool {
	p := &sync.Pool{
		New: connetcToService,
	}
	for i := 0; i < 10; i++ {
		p.Put(p.New)
	}
	return p
}

func startNetworkDaemon() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		connpool := warmServiceConnCache()

		server, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			log.Fatalf("Cannot listen: %v", err)
		}
		wg.Done()

		defer server.Close()
		for {
			conn, err := server.Accept()
			if err != nil {
				log.Printf("Cannot accept: %v", err)
				continue
			}
			// connetcToService()

			svcConn := connpool.Get()
			fmt.Fprintf(conn, "")
			connpool.Put(svcConn)

			fmt.Println("Accept a connection.")
			conn.Close()
		}
	}()
	return &wg
}

func init() {
	daemonStarted := startNetworkDaemon()
	daemonStarted.Wait()
}

// Benchmark〇〇で始まる関数は、ベンチマークテストとして認識される
// go test -benchtime=10s -bench=.
func BenchmarkNetworkRequest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			b.Fatalf("Cannot dial: %v", err)
		}
		if _, err := ioutil.ReadAll(conn); err != nil {
			b.Fatalf("Cannot read: %v", err)
		}
		conn.Close()
	}
}
