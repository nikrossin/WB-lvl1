package main

import (
	"fmt"
	"os"
)

func main() {
	arr := []int{2, 4, 5, 8, 10}
	ch := make(chan int)

	for _, val := range arr {
		go func(val int) {
			ch <- val * val
		}(val)
		fmt.Fprintln(os.Stdout, <-ch)
	}
}
