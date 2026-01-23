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
	nodeIndex := make(map[*Node_138]int)
	var newNodes []*Node_138
	for i, p := 0, head; p != nil; p = p.Next {
		nodeIndex[p] = i
		newNodes = append(newNodes, &Node_138{Val: p.Val, Random: p.Random})
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
