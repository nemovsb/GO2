package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
)

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	wg := sync.WaitGroup{}
	n := 10
	win := make(chan string, n)

	wg.Add(n)
	for i := 0; i <= n; i++ {

		go func(gouroutineID int) {
			for j := 0; j <= 1000000; j++ {
				//fmt.Printf("Gorutine %d, count = %d\n", gouroutineID, j)
				//time.Sleep(1 * time.Millisecond)
				if j%1000 == 0 {
					runtime.Gosched()
				}
			}
			win <- fmt.Sprintf("Gorutine %d win\n", gouroutineID)
			fmt.Printf("Gorutine %d done\n", gouroutineID)
			//close(done)
			wg.Done()
		}(i)

	}

	fmt.Println(<-win)
	wg.Wait()
}
