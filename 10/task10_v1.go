package main

import "fmt"

func main() {
	var nums = []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, -15, -17, -9, 1}
	var set = make(map[int][]float32)

	for _, val := range nums {

		key := int(val/10) * 10 //получаем круглые целые числа для определения группы
		set[key] = append(set[key], val)

	}
	fmt.Println(set)

}
