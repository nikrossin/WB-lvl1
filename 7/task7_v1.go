package main

import (
	"fmt"
	"sync"
)

func main() {
	var data = make(map[int]int)
	var wg sync.WaitGroup
	var mutex sync.Mutex
	wg.Add(10)

	for i := 0; i < 10; i++ {
		//конкурентная запись 10 значений в мапу
		go func(n int) {
			defer wg.Done()
			mutex.Lock() // блокируем мьютекс для записи данных в мапу (общему ресурсу)
			data[n] = n * n
			//fmt.Println(data[n])
			mutex.Unlock() //после записи - разблокируем
		}(i)
	}
	wg.Wait() //ожидаем завершение всех горутин

	mutex.Lock() //блокируем мьютекс
	for _, val := range data {
		fmt.Println(val)
	}
	mutex.Unlock() //разблокируем мьютекс
	fmt.Println("Done")

}
