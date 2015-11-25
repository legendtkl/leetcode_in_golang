//Solution:	Hash
//Status:	AC

package main

import (
	"fmt"
)

func TwoSum(nums []int, target int) (a, b int) {
	exist := make(map[int]int)
	n := len(nums)
	for i := 0; i != n; i++ {
		if exist[target-nums[i]] != 0 {
			a = exist[target-nums[i]]
			b = i + 1
		}
		exist[nums[i]] = i + 1
	}
	return
}

func main() {
	nums := []int{2, 7, 11, 13}
	fmt.Println(TwoSum(nums, 13))
}
