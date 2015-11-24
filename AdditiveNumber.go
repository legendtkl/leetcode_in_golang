//Solution: DFS
//Status:	AC

package main

import (
	"fmt"
	"strconv"
)

var ret bool

func DFS(num string, start, pos1, pos2 int) {
	n := len(num)

	if ret {
		return
	}
	if (pos1 > start && num[start] == '0') || (pos2 > pos1 + 1 && num[pos1+1] == '0') {
		return
	}
	if pos2 == n-1 {
		ret = true
		return
	}

	var sum string
	index1 := pos1
	index2 := pos2
	flag := 0

	for index1 >= start || index2 > pos1 || flag == 1 {
		local := flag
		if index1 >= start {
			local += int(num[index1] - '0')
			index1 --
		}
		if index2 > pos1 {
			local += int(num[index2] - '0')
			index2 --
		}
		flag = local / 10
		local %= 10
		sum = strconv.Itoa(local) + sum
	}

	l := len(sum)
	if pos2 + l < n && sum == num[pos2+1:pos2+1+l] {
		DFS(num, pos1+1, pos2, pos2+l)
	} 
}

func isAdditveNumber(num string) bool {
	n := len(num)
	for i := 0; i < n; i++ {
		for j := i+1; j < n-1; j++ {
			DFS(num, 0, i, j)
		}
	}
	return ret
}

func main() {
	num := "199100199"
	fmt.Println(isAdditveNumber(num))	
	//fmt.Println(ret)
}