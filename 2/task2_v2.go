package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	arr := []int{2, 4, 5, 8, 10}
	var wg sync.WaitGroup
	wg.Add(5)

	for _, val := range arr {
		go func(param int) {
			defer wg.Done()
			fmt.Fprintln(os.Stdout, param*param)
		}(val)
	}
	wg.Wait()
}
