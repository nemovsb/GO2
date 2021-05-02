package main

import (
	"fmt"
	"time"
)

func main() {

	done := make(chan string, 1)

	n := 2

	for i := 0; i <= n; i++ {

		go func(gouroutineID int) {
			for j := 0; j <= 1000; j++ {
				//fmt.Printf("Gorutine %d, count = %d\n", gouroutineID, j)
				time.Sleep(1 * time.Millisecond)
			}
			done <- fmt.Sprintf("Gorutine %d win", gouroutineID)
		}(i)

	}

	//<-done
	fmt.Println(<-done)

}
