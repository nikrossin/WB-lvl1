package main

import (
	"fmt"
	"strings"
)

func checkUnique(str string) bool {

	var symbols = make(map[rune]struct{})
	str = strings.ToLower(str) //приводим все символы в единый регистр
	//проход за n символов
	for _, val := range str {
		if _, ok := symbols[val]; ok { //если такой символ уже записан в мапу, возвращаем false
			return false
		} else {
			symbols[val] = struct{}{}
		}
	}
	return true //иначе строка уникальна
}
func main() {

	s := "abcdifghzxvV "
	fmt.Println(checkUnique(s))

}
