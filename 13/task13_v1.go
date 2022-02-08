package main

import "fmt"

func swap(a *int, b *int) {
	*a, *b = *b, *a
}

func main() {
	x := 1
	y := 2
	swap(&x, &y)
	fmt.Println(x, y)

}
