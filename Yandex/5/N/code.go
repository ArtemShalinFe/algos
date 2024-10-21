package main

// <template>
type Node struct {
	value int
	left  *Node
	right *Node
	size  int
}

// <template>

func split(node *Node, k int) (*Node, *Node) {
	// Базовый случай
	if node == nil {
		return nil, nil
	}

	// Получаем размер левого и правого поддерева
	ls, rs := size(node.left, node.right)

	// Если размер левого дерева + 1 больше чем индекс нужной вершины
	// то разделяем левое
	if ls+1 > k {
		ln, rn := split(node.left, k)
		_, rnSize := size(ln, rn)

		node.size = node.size - ls + rnSize
		node.left = rn

		return ln, node
	}

	// Если размер правого поддерева больше чем индекс нужной вершины
	// то разделяем правое
	ln, rn := split(node.right, k-(ls+1))
	lnSize, _ := size(ln, rn)

	node.size = node.size - rs + lnSize
	node.right = ln

	return node, rn
}

func size(l, r *Node) (ls, rs int) {
	if l != nil {
		ls = l.size
	}
	if r != nil {
		rs = r.size
	}

	return
}

func test() {
	node1 := &Node{3, nil, nil, 1}
	node2 := &Node{2, nil, node1, 2}
	node3 := &Node{8, nil, nil, 1}
	node4 := &Node{11, nil, nil, 1}
	node5 := &Node{10, node3, node4, 3}
	node6 := &Node{5, node2, node5, 6}
	left, right := split(node6, 4)
	if left.size != 4 {
		panic("WA")
	}
	if right.size != 2 {
		panic("WA")
	}
}
