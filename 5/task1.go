package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	wg := sync.WaitGroup{}
	n := 10
	win := make(chan string, n)

	wg.Add(n)
	for i := 0; i <= n; i++ {

		go func(gouroutineID int) {
			for j := 0; j <= 110; j++ {
				//fmt.Printf("Gorutine %d, count = %d\n", gouroutineID, j)
				time.Sleep(1 * time.Millisecond)
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
