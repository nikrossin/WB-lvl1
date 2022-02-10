package main

import "fmt"

func setBit(num int64, bitPos int, bitVal int) int64 {
	//в зависимости от значения бита - используем различные битовые операции
	if bitVal == 1 {
		return num | 1<<(bitPos-1) //побитовое сложение с маской установленного бита на нужной позиции
	} else {
		return num & ^(1 << (bitPos - 1)) //побитовое умножение с инвертированной маской установленного бита на нужной позиции
	}
}

func main() {

	var num int64
	var bitPos, bitVal int
	fmt.Scanln(&num, &bitPos, &bitVal) // ввод числа, позиция бита и его значение
	fmt.Printf("Before set - bin: %b  dec: %d\n", num, num)
	num = setBit(num, bitPos, bitVal)
	//для отрицательных чисел бинарный вывод в прямом коде
	fmt.Printf("After set  - bin: %b  dec: %d", num, num)

}
