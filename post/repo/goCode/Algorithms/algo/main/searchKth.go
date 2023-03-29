package main

/*一个带头节点的单链表，只给出头节点 head,
* 在不改变链表的情况下，设计尽可能高效的算法，查找链表中倒数第 k 个元素
* 成功则返回（输出）节点上的值，并返回 1
* 失败则返回 0
 */

import (
	. "fmt"
)

type Element interface{}

type Node struct {
	data Element
	next *Node
}

func NewNode(d Element) *Node {
	return &Node{data: d}
}

func (head *Node) Add(d Element) {
	p := head
	for p.next != nil {
		p = p.next
	}
	n := NewNode(d)
	p.next = n
}

func searchKth(k int, head *Node) (data Element, err int) {
	q, p := head.next, head.next
	count := 0

	// p 作为前置指针，会一直遍历到链表的尾部
	for p != nil {
		if count < k {
			// 计算 p 与 q 的距离，为要寻找的倒数 k
			count++
		} else {
			// 当 p 与 q 到达指定距离后，q 作为后置指针，也作为最终的答案指针，开始移动
			// 随着 p 到达链表尾部，q 也指向了最终的答案节点
			q = q.next
		}
		// 循环中，p 会一直往下走
		p = p.next
	}

	if count < k {
		err = 0
	} else {
		data = q.data
		err = 1
	}
	return
}

func main() {
	var d *Node
	d = NewNode("head")
	for i := 0; i < 10; i++ {
		d.Add(i + 1)
	}

	dd := d
	Println("elements in source list:")
	for dd != nil {
		Println(dd.data)
		dd = dd.next
	}

	ret, err := searchKth(8, d)

	if err != 0 {
		Printf("8th element is :%d", ret)
	}
}
