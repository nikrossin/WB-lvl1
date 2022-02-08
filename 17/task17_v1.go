package main

import "fmt"

func binSearch(arr []int, num int) int {
	l, r := -1, len(arr)

	for r-l > 1 {
		m := (l + r) / 2

		if num > arr[m] {
			l = m
		} else {
			r = m
		}
	}
	return r
}

func main() {

	arr := []int{1, 2, 2, 2, 2, 3, 5, 8, 9, 11}
	fmt.Println(binSearch(arr, 9))

}
