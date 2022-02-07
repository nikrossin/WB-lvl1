package main

import "fmt"

func setBit(num int64, bitPos int, bitVal int) int64 {
	if bitVal == 1 {
		return num | 1<<(bitPos-1)
	} else {
		return num & ^(1 << (bitPos - 1))
	}
}

func main() {

	var num int64
	var bitPos, bitVal int
	fmt.Scanln(&num, &bitPos, &bitVal)
	fmt.Printf("Before set - bin: %b  dec: %d\n", num, num)
	num = setBit(num, bitPos, bitVal)
	//для отрицательных чисел бинарный вывод в прямом коде
	fmt.Printf("After set  - bin: %b  dec: %d", num, num)

}
