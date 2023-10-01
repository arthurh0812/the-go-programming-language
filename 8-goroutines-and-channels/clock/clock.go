package main

import (
	"io"
	"log"
	"net"
	"sync"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}

	var wg sync.WaitGroup
	done := make(chan bool)

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Println(err)
				continue
			}
			wg.Add(1)
			go func() {
				handle(conn)
				wg.Done()
				done <- true
			}()
		}
	}()

	go func() {
		wg.Wait()
		close(done)
	}()

	var count int
	for range done {
		count++
	}
	log.Printf("Established %d connections\n", count)
}

func handle(conn net.Conn) {
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, time.Now().Format("00:00:00\n"))
		if err != nil {
			log.Println(err)
			return
		}
		time.Sleep(time.Second)
	}
}
