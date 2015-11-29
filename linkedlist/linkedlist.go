// single linked list implement
// go has package container/list: a doubly linked list

package linkedlist

import (
	//	"fmt"
	"reflect"
)

type Node struct {
	Value interface{}
	Next  *Node
}

func (n *Node) Init(v interface{}) *Node {
	n.Value = v
	n.Next = nil
	return n
}

func NewNode(v interface{}) *Node {
	return new(Node).Init(v)
}

func CreatList(nums interface{}) (*Node, bool) {
	v := reflect.ValueOf(nums)
	if v.Kind() != reflect.Slice {
		return nil, false
	}

	n := v.Len()
	if n == 0 {
		return nil, false
	}

	head := new(Node).Init(v.Index(0).Interface())
	pre := head
	for i := 1; i < n; i++ {
		pre.Next = new(Node).Init(v.Index(i).Interface())
		pre = pre.Next
	}
	/*
		for pre = head; pre != nil; pre = pre.Next {
			fmt.Println(pre.Value)
		}*/
	return head, true
}
