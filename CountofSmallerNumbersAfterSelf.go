package main

import (
	"fmt"
)

func lower_bound(nums []int, key int) int {
	if key < nums[0] {
		return 0
	}

	first := 0
	count := len(nums)

	for count > 0 {
		mid := first
		step := count / 2
		mid += step

		if nums[mid] < key {
			mid++
			first = mid
			count -= step + 1
		} else {
			count = step
		}
	}
	return first
}

func countSmaller(nums []int) []int {
	n := len(nums)
	if len(nums) == 0 {
		return []int{}
	}

	ret := make([]int, n)
	rightNums := make([]int, 0, n)

	for i := n - 1; i >= 0; i-- {
		if len(rightNums) == 0 {
			ret[i] = 0
			rightNums = append(rightNums, nums[i])
		} else {
			find := lower_bound(rightNums, nums[i])
			ret[i] = find

			rightNums = append(rightNums, 0)
			copy(rightNums[find+1:], rightNums[find:len(rightNums)])
			rightNums[find] = nums[i]
		}
	}
	return ret
}

func main() {
	A := []int{5, 2, 6, 1}
	fmt.Println(countSmaller(A))
}
