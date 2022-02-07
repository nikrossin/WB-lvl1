package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	var n int
	ch := make(chan int)
	var wg sync.WaitGroup

	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(num int, ch <-chan int) {
			for val := range ch {
				fmt.Println("id: ", num, "value ", val)
			}
			wg.Done()
		}(i, ch)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	func() {
		for {
			select {
			case <-stop:
				return
			default:
				ch <- int(rand.Intn(100))
			}
		}
	}()

	close(ch)
	wg.Wait()

}
