package main

// <template>
type ListNode struct {
	data string
	next *ListNode
}

// <template>

func Solution(head *ListNode, elem string) int {
	idx := 0
	for {
		if head.next == nil {
			return -1
		}
		if elem == head.data {
			return idx
		}
		idx++
		head = head.next
	}
}

func test() {
	node3 := ListNode{"node3", nil}
	node2 := ListNode{"node2", &node3}
	node1 := ListNode{"node1", &node2}
	node0 := ListNode{"node0", &node1}
	/*idx :=*/ Solution(&node0, "node2")
	// result is : idx == 2
}
