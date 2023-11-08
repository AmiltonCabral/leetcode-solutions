package main

import "fmt"

func removeDuplicates(nums []int) int {
	k := len(nums)
	j := 0
	actual := nums[0]
	rep := 0

	for i := 0; i < len(nums); i++ {
		nums[j] = nums[i]

		if nums[i] == actual {
			rep++
		} else {
			rep = 1
			actual = nums[i]
		}

		if rep < 3 {
			j++
		} else {
			k--
		}

	}

	return k
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	k := removeDuplicates(nums)
	fmt.Println(k, nums)
}
