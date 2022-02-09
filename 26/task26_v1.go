package main

import (
	"fmt"
	"strings"
)

func checkUnique(str string) bool {

	var symbols = make(map[rune]struct{})
	str = strings.ToLower(str)

	for _, val := range str {
		if _, ok := symbols[val]; ok {
			return false
		} else {
			symbols[val] = struct{}{}
		}
	}
	return true
}
func main() {

	s := "abcdifghzxvV "
	fmt.Println(checkUnique(s))

}
