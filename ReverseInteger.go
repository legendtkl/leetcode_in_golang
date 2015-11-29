// Solution:	Imitaion
// Status:		AC

package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	var ret int64 = 0
	for x != 0 {
		ret = ret*10 + int64(x%10)
		x /= 10
	}
	if ret > math.MaxInt32 || ret < math.MinInt32 {
		ret = 0
	}

	return int(ret)
}

func main() {
	x := -1234567893
	//	fmt.Println(math.MaxInt32)
	//	fmt.Println(math.MinInt32)
	fmt.Println(reverse(x))
}
