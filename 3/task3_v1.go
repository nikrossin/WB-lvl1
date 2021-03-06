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
		//главная горутина ожидает получения значения и складывает
		sum += <-ch
	}
	fmt.Fprintln(os.Stdout, sum)
}
