package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	newHead := &Node{head.Val, nil, head.Random}
	randMap := map[*Node]*Node{
		head: newHead,
	}
	p, q := head, newHead
	for p.Next != nil {
		randMap[p] = q
		p = p.Next
		q.Next = &Node{p.Val, p.Next, p.Random}
		randMap[p] = q.Next
		q = q.Next
	}
	q = newHead
	for q != nil {
		if q.Random != nil {
			q.Random = randMap[q.Random]
		}
		q = q.Next
	}
	return newHead
}
