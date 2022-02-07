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
		go func(n int) {
			defer wg.Done()
			mutex.Lock()
			data[n] = n * n
			//fmt.Println(data[n])
			mutex.Unlock()
		}(i)
	}
	wg.Wait()

	mutex.Lock()
	for _, val := range data {
		fmt.Println(val)
	}
	mutex.Unlock()
	fmt.Println("Done")

}
