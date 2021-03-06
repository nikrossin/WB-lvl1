package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	sync.RWMutex
	count int
}

//конструктор счетчика
func NewCounter() Counter {
	return Counter{}
}

//метод инкремента
func (c *Counter) Inc() {
	c.Lock() //блокируем ресурс
	defer c.Unlock()
	c.count++
}

//метод печати значения счетчика
func (c *Counter) Get() int {
	c.RLock() //блокируем ресурс на запись(читать можно)
	defer c.RUnlock()
	return c.count
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
