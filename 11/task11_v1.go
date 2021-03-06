package main

import "fmt"

func cross(arr1 []int, arr2 []int) []int {
	var cross = make(map[int]int)
	var arrCross []int
	//ключ - значение массива 1, значени мапы - количество значений в массиве 1
	for _, val := range arr1 {
		if _, ok := cross[val]; ok {
			cross[val]++
		} else {
			cross[val] = 1
		}
	}
	//если такое значение существует в мапе, декриментируем значение и записывает элемент как пересечение
	//если в первом массиве две 5, а во втором одна 5, в пересечение только одна пятерка
	for _, val := range arr2 {
		if _, ok := cross[val]; ok {
			if cross[val] != 0 {
				cross[val]--
				arrCross = append(arrCross, val)
			}
		}
	}
	return arrCross
}

func main() {
	nums1 := []int{2, 5, 5, 3, 1, 4, 5}
	nums2 := []int{5, 7, 6, 4, 8, 5, 3, 5, 5, 5}
	nums3 := cross(nums1, nums2)
	fmt.Println(nums3)

}
