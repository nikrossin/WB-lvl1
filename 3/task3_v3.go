package main

import (
	"fmt"
	"os"
)

func main() {
	arr := []int{2, 4, 5, 8, 10}
	ch := make(chan int)
	var sum int
	for _, val := range arr {
		//каждая горутина вычисляет значение квадрата
		go func(val int) {
			ch <- val * val
		}(val)

	}
	//главная горутина блокируется до получения значения в канале
	for i := 0; i < len(arr); i++ {
		sum += <-ch
	}
	fmt.Fprintln(os.Stdout, sum)
}
