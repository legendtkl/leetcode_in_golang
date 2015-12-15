package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func lengthOfLongestSubstring(s string) int {
	var hash [128]int
	start := -1
	n := len(s)
	ret := 0

	for i := 0; i < n; i++ {
		if hash[int(s[i])] != -1 {
			start = max(start, hash[int(s[i])])
		}
		ret = max(ret, i-start)
		hash[int(s[i])] = i
	}
	return ret
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abc"))
}
