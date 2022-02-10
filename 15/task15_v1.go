package main

import (
	"fmt"
)

var justString string

func createHugeString(length int) string {
	var str string
	for i := 0; i < length; i++ {
		str += "ы"
	}
	return str
}

func someFunc() {
	v := createHugeString(1 << 10)
	fmt.Println("Исходная строка:", cap([]byte(v)), len([]byte(v)))

	//строка - слайс байт, при срезе оценим len и capacity
	justString = v[:99]
	fmt.Println("Срез исходной строки", justString)
	fmt.Println(cap([]byte(justString)), len([]byte(justString)))

	justString = string([]rune(v)[:99])
	fmt.Println("Срез строки, преобразованной в руны", justString)
	fmt.Println(cap([]byte(justString)), len([]byte(justString)))

	strRune := make([]rune, 0, 99)
	strRune = append(strRune, []rune(v)[:99]...)
	fmt.Println("Срез строки с выделением cap", string(strRune))
	fmt.Println(cap(strRune), len(strRune))

	var strRune2 []rune
	strRune2 = []rune(v)[:99]
	fmt.Println(string(strRune2))
	fmt.Println(cap(strRune2), len(strRune2))
}

func main() {
	someFunc()
	fmt.Println(justString)
}
