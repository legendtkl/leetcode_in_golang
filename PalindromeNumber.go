package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	div := 1
	for x/div >= 10 {
		div *= 10
	}
	div1 := 1
	for div > div1 {
		head := x / div % 10
		tail := x / div1 % 10
		if head != tail {
			return false
		}
		div /= 10
		div1 *= 10
	}
	return true
}

func main() {
	nums := []int{-1, 0, 1212, 121}
	for i := 0; i < len(nums); i++ {
		println(isPalindrome(nums[i]))
	}
}
