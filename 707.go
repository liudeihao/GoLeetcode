package main

type Node_707 struct {
	Val  int
	Next *Node_707
}

type MyLinkedList struct {
	dummy *Node_707
	tail  *Node_707
	len   int
}

func Constructor_707() MyLinkedList {
	dummy := &Node_707{}

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
	newNode := &Node_707{val, this.dummy.Next}
	this.dummy.Next = newNode
	if newNode.Next == nil {
		this.tail = newNode
	}
	this.len++
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.tail.Next = &Node_707{val, nil}
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
	newNode := &Node_707{val, p.Next}
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
