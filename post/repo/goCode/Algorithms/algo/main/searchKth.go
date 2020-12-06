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

type node struct{
    data Element
    next *node
}

func NewNode(d Element) *node{
    return &node{data:d}
}

func (head *node) Add(d Element) {
    p := head
    for p.next != nil {
        p = p.next
    }
    n := NewNode(d)
    p.next = n
}

func searchKth(k int,head *node) (data Element,err int){
    q,p := head.next,head.next
    count := 0

    for p != nil {
        if count < k {
            count++
        }else{
            q = q.next
        }
        p = p.next
    }

    if count < k {
        err = 0
    }else{
        data = q.data
        err = 1
    }
    return
}

func main() {
    var d *node
    d = NewNode("head")
    for i := 0; i < 10; i++ {
        d.Add(i+1)
    }

    dd := d
    Println("elements in source list:")
    for dd != nil {
        Println(dd.data)
        dd = dd.next
    }

    ret,err := searchKth(8,d)

    if err != 0 {
        Printf("8th element is :%d",ret)
    }
}