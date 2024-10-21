package main

// <template>
type Node struct {
	value int
	left  *Node
	right *Node
}

// <template>

func Solution(root *Node) bool {
	if root == nil {
		return true
	}

	l := height(root.left)
	r := height(root.right)

	if max(l, r)-min(l, r) <= 1 && Solution(root.left) && Solution(root.right) {
		return true
	}

	return false
}

func height(root *Node) int {
	if root == nil {
		return 0
	}

	l := height(root.left)
	r := height(root.right)

	return 1 + max(l, r)
}

func test() {
	node1 := Node{1, nil, nil}
	node2 := Node{-5, nil, nil}
	node3 := Node{3, &node1, &node2}
	node4 := Node{10, nil, nil}
	node5 := Node{2, &node3, &node4}
	if !Solution(&node5) {
		panic("WA")
	}
}

/*
			2
		3		10
	1		-5
*/
