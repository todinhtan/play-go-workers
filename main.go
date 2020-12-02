package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/todinhtan/play-go-workers/worker"
)

func receive(rc <-chan string, conn net.Conn) {
	for {
		select {
		case r := <-rc:
			fmt.Fprintln(conn, "Result:", r)
		default:
			return
		}
	}
}

var wg sync.WaitGroup

var numWorkers int

func main() {
	args := os.Args[1:]
	if len(args) <= 0 {
		fmt.Println("Please initialize max workers")
		return
	}

	s := args[0]
	numWorkers, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Please provide a number")
		return
	}

	// init tcp server
	li, err := net.Listen("tcp", ":8080")
	defer li.Close()
	if err != nil {
		log.Fatalln(err)
		return
	}

	for {
		// read connection
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Fprintln(conn, "Please add some tasks")
		go handle(conn, numWorkers)
	}
}

func handle(conn net.Conn, nw int) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		txt := scanner.Text()
		ss := strings.Split(txt, ",")

		tl := len(ss)
		tasks := make(chan string, tl)
		results := make(chan string, tl)

		// register workers
		for i := 1; i <= nw; i++ {
			go worker.Worker(i, &wg, tasks, results)
		}

		// add item to tasks channel
		for _, s := range ss {
			wg.Add(1)
			tasks <- s
		}

		// wait for all workers finish their tasks
		wg.Wait()
		// print result to client
		receive(results, conn)
	}

	defer conn.Close()
}
