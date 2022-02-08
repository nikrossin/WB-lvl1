package main

import "fmt"

type Counter struct {
	count int
}

func NewCounter() Counter {
	return Counter{0}
}

func (c *Counter) Inc() {
	c.count++
}

func (c *Counter) PrintCounter() {
	fmt.Println("Counter:", c.count)
}
func main() {
	counter := NewCounter()
	ch := make(chan struct{})

	go func(c *Counter) {
		for i := 0; i < 20; i++ {
			c.Inc()
		}
		close(ch)
	}(&counter)
	<-ch
	counter.PrintCounter()

}
