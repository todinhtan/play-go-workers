package worker

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// sleep time
const st = 2 * time.Second

// Worker does the job
// receive identifer s, tasks t, return results r
func Worker(id int, wg *sync.WaitGroup, t <-chan string, r chan<- string) {
	for {
		select {
		case c := <-t:
			fmt.Printf("Worker %d started a new job \"%s\".\n", id, c)
			time.Sleep(st)
			r <- strings.ToUpper(c)
			fmt.Printf("Worker %d finished the job \"%s\".\n", id, c)
			wg.Done()
		default:
			return
		}
	}
}
