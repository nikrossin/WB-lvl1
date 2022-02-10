package main

import (
	"fmt"
	"sync"
	"time"
)

const sec = 1

// функция записи данных в канал
func write(ch chan<- int, stop <-chan struct{}) {
	var i int
	for {
		select {
		case <-stop: //инидкатор завершения работы горутины
			return
		default:
			ch <- i
			i += 10
		}
	}
}

//чтение из канала
func read(ch <-chan int, wg *sync.WaitGroup) {
	//читаем пока не закроют канал
	for val := range ch {
		fmt.Println(val)
	}
	wg.Done()

}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	ch := make(chan int)
	stop := make(chan struct{})

	go write(ch, stop)
	go read(ch, &wg)
	//работаем n секунд
	time.Sleep(time.Second * sec)
	//отправляем сигнал завершение горутины на запись
	stop <- struct{}{}
	close(ch) //закрываем канал
	wg.Wait() //используем ожидание для горутины на чтение, чтобы данные успели считаться из канала и вывестись в stdout
}
