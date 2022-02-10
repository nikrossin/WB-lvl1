package main

import (
	"fmt"
	"os"
)

func main() {
	arr := []int{2, 4, 5, 8, 10}
	ch := make(chan int)

	//анонимные функции вызываются в отдельных горутинах, для вычисления квадрата каждого элемента слайса
	for _, val := range arr {
		go func(val int) {
			ch <- val * val
		}(val)

	}
	//главная горутина блокируется до получения значения в канале
	for i := 0; i < len(arr); i++ {
		fmt.Fprintln(os.Stdout, <-ch)
	}
}
