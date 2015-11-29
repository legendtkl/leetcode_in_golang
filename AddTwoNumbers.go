package main

import (
	"./linkedlist"
	"fmt"
	//	"reflect"
)

func addTwoNumbers(L1, L2 *linkedlist.Node) *linkedlist.Node {
	flag := 0
	ret := linkedlist.NewNode(-1)
	p := ret

	for L1 != nil || L2 != nil || flag != 0 {
		sum := flag
		if L1 != nil {
			sum += L1.Value.(int)
			L1 = L1.Next
		}
		if L2 != nil {
			sum += L2.Value.(int)
			L2 = L2.Next
		}
		flag = sum / 10
		sum %= 10
		p.Next = linkedlist.NewNode(sum)
		p = p.Next
	}
	ret = ret.Next
	return ret
}
func main() {

	A := []int{2, 4, 3}
	B := []int{5, 6, 4}
	L1, _ := linkedlist.CreatList(A)
	L2, _ := linkedlist.CreatList(B)
	head := addTwoNumbers(L1, L2)
	for p := head; p != nil; p = p.Next {
		fmt.Println(p.Value)
	}
	/*
		root * Node = CreatList(A)
		for root != nil {
			fmt.Println(root.Value)
			root = root.Next
		}*/
}
