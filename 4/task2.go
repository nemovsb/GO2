package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/* Написать программу, которая при получении в канал сигнала SIGTERM останавливается не
позднее, чем за одну секунду (установить таймаут). */

func main() {

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Printf("Signal: %s\n", sig)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		done <- true

		<-ctx.Done()
		cancel()

		fmt.Println("exiting from gorutine (Timeout)")
		os.Exit(0)
	}()

	fmt.Println("awaiting signal SIGTERM")
	<-done
	fmt.Println("exiting from main")
}
