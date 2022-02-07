package main

import "fmt"

func main() {

	var ch1 = make(chan int)
	var ch2 = make(chan int)
	var sync = make(chan struct{})
	var arr = [5]int{1, 2, 3, 4, 5}

	go func() {
		for c := range ch1 {
			ch2 <- c * 2
		}
		close(ch2)
	}()

	go func() {
		for v := range ch2 {
			fmt.Println(v)
		}
		close(sync)
	}()

	for _, val := range arr {
		ch1 <- val
	}
	close(ch1)
	<-sync

}
