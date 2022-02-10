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
	//вводим кол-во воркеров
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		wg.Add(1) //добавляем горутину в группу
		go func(num int, ch <-chan int) {
			//читаем данные, пока не закроется канал
			for val := range ch {
				fmt.Println("id: ", num, "value ", val)
			}
			wg.Done()
		}(i, ch)
	}

	stop := make(chan os.Signal, 1) // канал системного сигнала
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	func() {
		for {
			select {
			case <-stop: // как только придет системный сигнла (ctl+c) - завершение функции
				return
			default:
				ch <- rand.Intn(100) //пишем данные
			}
		}
	}()
	close(stop)
	close(ch) // закрываем канал
	wg.Wait() // ждем, чтобы горутина успела считать данные из закрытого канала и вывела в stdout
	fmt.Println("PROGRAM STOP")

}
