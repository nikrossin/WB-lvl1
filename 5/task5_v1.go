package main

import (
	"fmt"
	"sync"
	"time"
)

const sec = 1

func write(ch chan<- int, stop <-chan bool) {
	var i int
	for {
		select {
		case <-stop:
			return
		default:
			ch <- i
			i += 10
		}
	}
}
func read(ch <-chan int, wg *sync.WaitGroup) {

	for val := range ch {
		fmt.Println(val)
	}
	wg.Done()

}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	ch := make(chan int)
	stop := make(chan bool)

	go write(ch, stop)
	go read(ch, &wg)

	time.Sleep(time.Second * sec)

	stop <- true
	close(ch)
	wg.Wait()
}
