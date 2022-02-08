package main

import (
	"fmt"
)

func typeOfVal(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int:", x)
	case string:
		fmt.Println("string:", x)
	case bool:
		fmt.Println("bool:", x)
	case chan int:
		fmt.Println("chan int:", x)
	default:
		fmt.Println("unknown")
	}
}

func main() {
	x1 := 2
	x2 := "hello"
	x3 := true
	x4 := make(chan int)

	typeOfVal(x1)
	typeOfVal(x2)
	typeOfVal(x3)
	typeOfVal(x4)
}
