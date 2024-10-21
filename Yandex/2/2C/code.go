package main

// <template>
type ListNode struct {
	data string
	next *ListNode
}

// <template>

func Solution(head *ListNode, idx int) *ListNode {
	if idx == 0 {
		return head.next
	} else {
		idx--
		head.next = Solution(head.next, idx)
	}

	return head
}

func test() {
	node3 := ListNode{"node3", nil}
	node2 := ListNode{"node2", &node3}
	node1 := ListNode{"node1", &node2}
	node0 := ListNode{"node0", &node1}
	/*newHead :=*/ Solution(&node0, 1)
	// result is : node0 -> node2 -> node3
}
