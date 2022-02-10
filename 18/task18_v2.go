package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	count int64
}

//конструктор счетчика
func NewCounter() Counter {
	return Counter{}
}

//метод инкремента
func (c *Counter) Inc() {
	atomic.AddInt64(&c.count, 1) //используем атомарную операцию увеличения, которая быстрее мьютексов
}

//метод печати значения счетчика
func (c *Counter) Get() int64 {
	return atomic.LoadInt64(&c.count) //атомарная операция чтения
}
func main() {
	counter := NewCounter()
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(c *Counter) {
			c.Inc()
			wg.Done()
		}(&counter)
	}
	wg.Wait()
	fmt.Println(counter.Get())

}
