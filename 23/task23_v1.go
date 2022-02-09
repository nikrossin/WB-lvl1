package main

import "fmt"

func deleteElem(arr []int, i int) []int {
	arr = append(arr[:i], arr[i+1:]...)
	return arr
}

func main() {
	var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	i := 0
	nums = deleteElem(nums, i)
	fmt.Println(nums)

}
