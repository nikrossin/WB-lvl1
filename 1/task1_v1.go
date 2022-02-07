package main

import "fmt"

type Human struct {
	firstName string
	lastName  string
	age       int
	country   string
}

func (h *Human) setInfo(fName string, lName string, age int, country string) {
	h.firstName, h.lastName, h.age, h.country = fName, lName, age, country
}
func (h Human) printInfo() {
	fmt.Printf("Name: %s %s, Age: %d, Country: %s", h.firstName, h.lastName, h.age, h.country)
}

type Action struct {
	id int
	Human
}

func main() {
	action := Action{}
	action.setInfo("Nikita", "Bulchuk", 22, "Russia")
	action.printInfo()
}
