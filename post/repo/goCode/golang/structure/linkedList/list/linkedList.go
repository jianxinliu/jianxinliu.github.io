package list

import (
	"errors"
	"fmt"
	// "strings"
	"strconv"
)

type linkedList struct{
	head *Node
	len int
	tail *Node
}


// =============================== list basic operations ====================================== //

func initList() *linkedList {
	n := &Node{0,nil}
	return &linkedList{n,0,n}
}


func GetList() *linkedList {
	return initList()
}

func (linkedList)NewNode(value int) *Node {
	return &Node{value,nil}
}

func (l *linkedList)Append(value int) *Node {
	n :=  l.NewNode(value)
	l.tail.Next = n
	l.tail = n
	l.len += 1
	return l.head
}

func (l *linkedList)PreAppend(value int) *Node {
	n := &Node{value,l.head.Next}
	l.head.Next = n
	l.len += 1
	return l.head
}

func (l *linkedList)Size() int {
	return l.len
}

// remove remove the specific node
// return error if the arguments value not exist
func (l *linkedList)Remove(value int) error {
	n ,p := l.head,l.head
	if l.tail.Value == value {
		l.RemoveTail()
	}else{
		var i = 1
		for ; n.Next != nil;i++ {
			if n.Value == value {
				break
			}
			p = n
			n = n.Next
		}
		if i > l.len {
			return errors.New("dont have the value:"+strconv.Itoa(value))
		}else{
			p.Next = n.Next
			n = nil
			l.len -= 1
		}
	}
	return nil
}

func (l *linkedList)RemoveTail()  {
	n ,p:= l.head,l.head
	for n.Next != nil {
		p = n
		n = n.Next
	}
	l.len -= 1
	p.Next = nil
	n = nil
}

func (l *linkedList)Print()  {
	arr := l.ToArray()
	for v := range arr {
		fmt.Printf("%d->",arr[v])
	}
	fmt.Println()
}

func (l *linkedList)ToArray() []int {
	ret := make([]int,0)
	n := l.head.Next
	for n != nil{
		ret = append(ret,n.Value)
		n = n.Next
	}
	return ret
}

func Create(arr []int) *linkedList {
	len := len(arr)
	list := GetList()
	for i := 0; i < len; i++ {
		list.Append(arr[i])
	}
	return list
}

// ================================== list advanced operations =============================== //

// reverse1 reverse the linkedlist by preappend
func (l *linkedList)Reverse1() *linkedList {
	list := GetList()
	n := l.head
	for n != nil{
		list.PreAppend(n.Value)
		n = n.Next
	}
	return list
}

// reverse2 reverse the linkedlist by Recursive
// func (l *linkedList)Reverse2() *linkedList {
// 	// head := &Node{0,nil}
	
// }

// reverseInstant reverse the linkedlist itself,not return a new list
func (l *linkedList)ReverseInstant() *Node {
	n := l.head

    l.head = nil
    for n != nil {
        m := n
        n = n.Next
        m.Next = l.head
        l.head = m
    }
    return l.head
}

// func (l *linkedList)reverse() *Node {
// 	return &Node{0,nil}
// }