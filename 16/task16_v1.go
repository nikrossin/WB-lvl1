package main

import "fmt"

func quickSort(arr []int, l int, r int) {
	v := arr[(l+r)/2]
	i, j := l, r

	for i <= j {
		for arr[i] < v {
			i++
		}
		for arr[j] > v {
			j--
		}
		if i <= j {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
		if i < r {
			quickSort(arr, i, r)
		}
		if l < j {
			quickSort(arr, l, j)
		}
	}

}

func main() {

	arr := []int{5, 3, 8, 1, 2, 1, 1, 1, 1, 1, 9, 0, 76, 3}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
