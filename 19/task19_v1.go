package main

import "fmt"

func stringReverse(str string) string {
	symbols := []rune(str)

	for i := 0; i < len(symbols)/2; i++ {
		symbols[i], symbols[len(symbols)-1-i] = symbols[len(symbols)-1-i], symbols[i]
	}
	return string(symbols)
}

func main() {
	var s = "HeLlo world, 世 界! Привет мир"
	fmt.Println(stringReverse(s))
}
