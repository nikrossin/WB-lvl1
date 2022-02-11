package main

import "fmt"

func update(p *int) {
	fmt.Println(*p)
	b := 2
	p = &b
	fmt.Println(*p)
}

func main() {
	var (
		a = 1
		p = &a
	)
	fmt.Println(*p)
	update(p)
	fmt.Println(*p)
}
