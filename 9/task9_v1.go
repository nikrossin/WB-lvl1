package main

import "fmt"

func main() {

	var ch1 = make(chan int)
	var ch2 = make(chan int)
	var sync = make(chan struct{})
	var arr = [5]int{1, 2, 3, 4, 5}

	//читаем данные из 1 канала, новое значение пишем во 2 канал
	go func() {
		for c := range ch1 {
			ch2 <- c * 2
		}
		close(ch2) //после закрытия 1 канала, горутина завершается и закрывает 2 канал
	}()

	//читаем данные из канала, пока его не закроют
	go func() {
		for v := range ch2 {
			fmt.Println(v)
		}
		close(sync) //канал синхронизации с главной горутиной
	}()

	// главная горутина
	for _, val := range arr {
		ch1 <- val
	}
	close(ch1) //после записи всез значений	- закрываем канал
	<-sync     // ждем завершения горутины, которая выводит данные

}
