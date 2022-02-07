package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	arr := []int{2, 4, 5, 8, 10}
	var m sync.Mutex
	var wg sync.WaitGroup
	var sum int

	for _, val := range arr {

		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			val = val * val
			m.Lock()
			sum += val
			m.Unlock()
		}(val)
	}
	wg.Wait()
	fmt.Fprintln(os.Stdout, sum)
}
