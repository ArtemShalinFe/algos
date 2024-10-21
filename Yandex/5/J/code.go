package main

type Node struct {
	value int
	left  *Node
	right *Node
}

func insert(root *Node, key int) *Node {
	if root == nil {
		return &Node{value: key}
	}

	if key < root.value {
		root.left = insert(root.left, key)
	} else if key >= root.value {
		root.right = insert(root.right, key)
	}

	return root
}

func test() {
	node1 := Node{7, nil, nil}
	node2 := Node{8, &node1, nil}
	node3 := Node{7, nil, &node2}
	newHead := insert(&node3, 6)
	if newHead != &node3 {
		panic("WA")
	}
	if newHead.left.value != 6 {
		panic("WA")
	}
}
