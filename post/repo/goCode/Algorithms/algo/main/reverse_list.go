package main

/**
1. 结构体中不允许递归定义，需要时应该用指针
2. 不能通过指针操作结构体（Golang不支持指针移动操作，如 node.Data = 2）
*/

import (
	. "fmt"
)

type Element interface{}

type node struct {
	Data Element
	Next *node
}

func NewNode(data Element) *node {
	return &node{Data: data}
}

func reverse_list(head *node) *node {
	n := head

	head = nil
	for n != nil {
		m := n
		n = n.Next
		m.Next = head
		head = m
	}
	return head
}

func (head *node) Add(data Element) {
	p := head
	for p.Next != nil {
		p = p.Next
	}
	n := NewNode(data)
	p.Next = n
}

func main() {
	var d *node
	d = NewNode("head")
	for i := 0; i < 8; i++ {
		d.Add(i + 1)
	}

	ret := reverse_list(d)
	for ret != nil && ret.Data != "head" {
		Println(ret.Data)
		ret = ret.Next
	}
}
