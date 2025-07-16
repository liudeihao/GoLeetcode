package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type MyLinkedList struct {
	dummy *Node
	tail  *Node
	len   int
}

func Constructor() MyLinkedList {
	dummy := &Node{}

	return MyLinkedList{
		dummy: dummy,
		tail:  dummy,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.len {
		return -1
	}
	p := this.dummy
	for index >= 0 {
		p = p.Next
		index--
	}
	return p.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &Node{val, this.dummy.Next}
	this.dummy.Next = newNode
	if newNode.Next == nil {
		this.tail = newNode
	}
	this.len++
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.tail.Next = &Node{val, nil}
	this.tail = this.tail.Next
	this.len++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.len {
		return
	}
	if index == this.len {
		this.AddAtTail(val)
		return
	}
	p := this.dummy
	for index > 0 {
		p = p.Next
		index--
	}
	newNode := &Node{val, p.Next}
	p.Next = newNode
	this.len++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.len {
		return
	}
	p := this.dummy
	for index > 0 {
		p = p.Next
		index--
	}
	if p.Next.Next == nil {
		this.tail = p
	}
	p.Next = p.Next.Next
	this.len--
}

func main() {
	obj := Constructor()
	fmt.Println(obj.Get(0))

	obj.AddAtHead(10)
	fmt.Println(obj.Get(0))

	obj.AddAtTail(20)
	fmt.Println(obj.Get(0))
	fmt.Println(obj.Get(1))

	obj.AddAtIndex(0, 50)
	obj.DeleteAtIndex(2)
	fmt.Println(obj.Get(0))
	fmt.Println(obj.Get(1))

}
