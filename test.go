package main

import "fmt"

func main() {
	var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	var nums2 []int
	fmt.Println(cap(nums), len(nums))
	nums2 = nums[2:6]
	nums[2] = 5
	fmt.Println(nums2)
	fmt.Println(cap(nums2), len(nums2))
}
