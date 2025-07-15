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
	nodeIndex := make(map[*Node]int)
	var newNodes []*Node
	for i, p := 0, head; p != nil; p = p.Next {
		nodeIndex[p] = i
		newNodes = append(newNodes, &Node{Val: p.Val, Random: p.Random})
		if i > 0 {
			newNodes[i-1].Next = newNodes[i]
		}
		i++
	}
	for i := 0; i < len(newNodes); i++ {
		idx, ok := nodeIndex[newNodes[i].Random]
		if ok {
			newNodes[i].Random = newNodes[idx]
		}
	}
	return newNodes[0]
}
