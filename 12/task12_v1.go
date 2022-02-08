package main

import (
	"fmt"
	"log"
)

type MySet struct {
	slice []string
}

func (s *MySet) Append(elem string) {
	s.slice = append(s.slice, elem)
}

func NewSet() MySet {
	return MySet{[]string{}}
}

func (s *MySet) Delete(elem string) error {
	for index, val := range s.slice {
		if val == elem {
			s.slice = append(s.slice[:index], s.slice[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("ERROR: Element '%s' not found!", elem)

}
func main() {

	var arr = []string{
		"cat",
		"cat",
		"dog",
		"tree",
		"dog",
		"世界万万",
		"рыба",
		"⌘",
		"рыба",
	}

	mySet := NewSet()
	for _, val := range arr {
		mySet.Append(val)
	}
	fmt.Println(mySet)
	err := mySet.Delete("kek")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(mySet)
}
