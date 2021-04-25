package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
)

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	wg := sync.WaitGroup{}
	n := 10
	wg.Add(n)

	var count uint32
	var myMu sync.Mutex
	for i := 0; i < n; i++ {

		go func(gouroutineID int) {
			myMu.Lock()
			defer myMu.Unlock()

			for j := 0; j < 10; j++ {
			}

			fmt.Printf("Gorutine %d done\n", gouroutineID)
			wg.Done()
		}(i)

	}
	wg.Wait()
	fmt.Printf("Count : %d\n", count)
}
