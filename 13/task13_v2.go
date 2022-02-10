package main

import "fmt"

func swap(a *int, b *int) {
	//используя математические вычисления
	*a = *a + *b
	*b = *a - *b
	*a = *a - *b

}

func main() {
	x := -10
	y := 2
	swap(&x, &y)
	fmt.Println(x, y)

}
