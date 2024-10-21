package main

import (
	"bufio"
	"os"
	"strconv"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

func printRange(root *Node, left int, right int) {
	if root.left != nil {
		printRange(root.left, left, right)
	}

	if root.value >= left && root.value <= right {
		writer := bufio.NewWriter(os.Stdout)
		writer.WriteString(strconv.Itoa(root.value))
		writer.WriteString(" ")
		writer.Flush()
	}

	if root.right != nil {
		printRange(root.right, left, right)
	}
}

func test() {
	node1 := Node{2, nil, nil}
	node2 := Node{1, nil, &node1}
	node3 := Node{8, nil, nil}
	node4 := Node{8, nil, &node3}
	node5 := Node{9, &node4, nil}
	node6 := Node{10, &node5, nil}
	node7 := Node{5, &node2, &node6}
	printRange(&node7, 2, 8)
	// expected output: 2 5 8 8
}
