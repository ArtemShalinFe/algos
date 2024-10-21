package main

import (
	"log"
)

// <template>
type Node struct {
	value int
	left  *Node
	right *Node
}

// <template>

func Solution(root *Node) bool {
	const MaxUint = ^uint(0)
	const MaxInt = int(MaxUint >> 1)
	const MinInt = -MaxInt - 1

	return solution(root, MinInt, MaxInt)
}

func solution(root *Node, left int, right int) bool {
	if root == nil {
		return true
	}

	if root.value <= left || root.value >= right {
		return false
	}

	leftOk := solution(root.left, left, root.value)
	rightOk := solution(root.right, root.value, right)

	return leftOk && rightOk
}

func test() {
	node1 := Node{1, nil, nil}
	node2 := Node{4, nil, nil}
	node3 := Node{3, &node1, &node2}
	node4 := Node{8, nil, nil}
	node5 := Node{5, &node3, &node4}
	if !Solution(&node5) {
		log.Fatal("no ok - WA")
	}
	node2.value = 5
	if Solution(&node5) {
		log.Fatal("ok - WA")
	}
}
