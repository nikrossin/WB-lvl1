package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	arr := []int{2, 4, 5, 8, 10}
	var wg sync.WaitGroup
	//5 горутин ожидаем
	wg.Add(5)

	for _, val := range arr {
		//каждая горутина вычисляет значение и выводит в stdout
		go func(param int) {
			defer wg.Done()
			fmt.Fprintln(os.Stdout, param*param)
		}(val)
	}
	//ожидаем завершения всех горутин
	wg.Wait()
}
