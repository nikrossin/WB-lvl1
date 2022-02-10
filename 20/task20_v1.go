package main

import (
	"fmt"
	"strings"
)

func wordsReverse(str string) string {
	words := strings.Split(str, " ") //преобразовываем в слайс строк

	for i := 0; i < len(words)/2; i++ {
		words[i], words[len(words)-1-i] = words[len(words)-1-i], words[i]
	}
	return strings.Join(words, " ") //получаем обратно строку
}

func main() {

	var s = "abs ddfe aa erer кука 世 界"

	fmt.Println(wordsReverse(s))
}
