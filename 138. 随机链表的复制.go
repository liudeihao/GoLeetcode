package main

type Node_138 struct {
	Val    int
	Next   *Node_138
	Random *Node_138
}

func copyRandomList(head *Node_138) *Node_138 {
	if head == nil {
		return nil
	}
	newHead := &Node_138{head.Val, nil, head.Random}
	randMap := map[*Node_138]*Node_138{
		head: newHead,
	}
	p, q := head, newHead
	for p.Next != nil {
		randMap[p] = q
		p = p.Next
		q.Next = &Node_138{p.Val, p.Next, p.Random}
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
