package main

import "fmt"

type Test struct {
	id *int
}

func main() {
	p := new([]int)
	a := make([]int, 0)
	var b []int
	fmt.Println(cap(*p), len(*p), cap(a), len(a), cap(b), len(b))
}
