package main

import (
	"fmt"
	"time"
)

var flag1 = make(chan struct{})
var flag2 = make(chan struct{})
var flag3 = make(chan struct{})

// завершение горутины по сигналу из отдельного канала (аналогично можно проверять значение из основного канала)
func worker1(ch <-chan int, exit <-chan int) {
	var n int

	for {

		select {
		case n = <-ch:
			fmt.Println("---------val: ", n, "------------")
		case <-exit:
			fmt.Println("exit gorutine 1")
			close(flag1)
			return
		default:
			fmt.Println("working gor 1")
		}
	}
}

//завершение горутины по значению переменной, переданной по указателю
func worker2(exit *bool) {

	for {
		if *exit {
			fmt.Println("exit gorutine 2")
			close(flag2)
			return
		}
		fmt.Println("working gor 2")

	}
}

// завершение по закрытию канала
func worker3(ch <-chan int) {
	for val := range ch {

		fmt.Println("working gor 3 -- val: ", val)
	}
	fmt.Println("exit gorutine 3")
	close(flag3)
}

func main() {
	ch := make(chan int)
	exit := make(chan int)
	exit2 := false

	go worker1(ch, exit)
	for i := 0; i < 10; i++ {
		time.Sleep(time.Nanosecond)
		ch <- i
	}
	exit <- 1
	<-flag1 // ожидание завершения горутины

	go worker2(&exit2)
	time.Sleep(time.Microsecond)
	exit2 = true
	<-flag2

	go worker3(ch)
	for i := 0; i < 10; i++ {
		time.Sleep(time.Nanosecond)
		ch <- i
	}
	close(ch)
	<-flag3

}
