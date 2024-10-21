/*
https://contest.yandex.ru/contest/24810/run-report/116000828/

### Принцип работы

	Алгоритм удаления узла у BST сначала выполняет поиск нужного узла по ключу.
	Если такой узел есть и:
	- у него нет подчиненных узлов, то узел просто удаляется.
	- у него есть только один подчиненный узел, то программа меняет родителя на подчиненный элемент
	- у него есть два подчиненных узла, то программа должна указать родителя для этих узлов. В текущей имплементации
	алгоритм выполняет поиск родителя в правом подчиненном узле и получает у него "самый левый элемент" - это самое маленькое
	значение из правого поддерева, и оно точно не меньше, чем любое значение левого поддерева.

### Доказательство корректности

	Перед удалением узла у BST алгоритм рекурсивно выполняет поиск нужного узла по ключу.

	Если узла с таким ключем нет, то алгоритм не изменяет дерево и возвращает корень не измененного дерева, что и требуется
	по условию задачи.

	Если такой узел есть и у него нет подчиненных узлов, то алгоритм удаляет найденный узел и возвращает корень дерева,
	структура которого принципиально не изменялась.

	Если такой узел есть и у него есть только один подчиненный узел, то алгоритм указывает в значении родителя значение
	подчиненного элемента и удаляет сам подчиненный элемент. Таким образом, алгоритм удаляет из дерева ключ и лишний узел,
	но сохраняет корректность BST потому, что после удаления дерево не распадается на 2 поддерева, а в вершине
	найденного узла установлен корректный для BST ключ.

	Если есть два подчиненных узла, то алгоритм выполняет поиск нового родителя в правом поддереве и запоминает из него
	самое маленькое значение, удалив при этом узел с самым маленьким значением. Так как запомненное значение точно не меньше,
	чем любое значение из левого поддерева, алгоритм указывает это значение как новое значение узла. Таким образом, алгоритм
	удаляет из дерева ключ и лишний узел, но сохраняет корректность BST потому, что после удаления узла дерево не распадается
	на 2 поддерева, а в вершине найденного узла установлен корректный для BST ключ.

### Временная сложность

	В худшем случае, чтобы удалить узел дерева, программе нужно будет обойти в глубину дерево состоящее из H узлов.

	Итого:
	- Временная сложность - О(H)

### Пространственная сложность

	В памяти программы хранятся все узлы дерева в количестве N единиц, которые связаны между собой.

	Итого:
	- Пространственная сложность - O(N)

*/

package main

type Node struct {
	value int
	left  *Node
	right *Node
}

func remove(node *Node, key int) *Node {
	// Базовый случай
	if node == nil {
		return node
	}

	switch {
	// Если значение в узле больше ключа - идем в левое поддерево
	case node.value > key:
		node.left = remove(node.left, key)
		// Если значение в узле меньше ключа - идем в правое поддерево
	case node.value < key:
		node.right = remove(node.right, key)
		// Значение равно значению ключа
	default:
		// Проверяем, вдруг у ключа только 1 подчиненный узел
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		}

		// Подчиненных узла точно два, поэтому
		// ищем самый левый узел в правом поддереве
		minNode := minValueOnLeftSideTree(node.right)

		// Обновляем значение в узле
		node.value = minNode.value

		// Удаляем самое левое значение в правом поддереве
		node.right = remove(node.right, minNode.value)
	}

	return node
}

func minValueOnLeftSideTree(node *Node) *Node {
	for node.left != nil {
		node = node.left
	}

	return node
}

/*
        5
    /       \
   1         10
    \       /
     3     8
    /     /
   2     6
*/

func test() {
	node1 := Node{2, nil, nil}
	node2 := Node{3, &node1, nil}
	node3 := Node{1, nil, &node2}
	node4 := Node{6, nil, nil}
	node5 := Node{8, &node4, nil}
	node6 := Node{10, &node5, nil}
	node7 := Node{5, &node3, &node6}
	newHead := remove(&node7, 10)
	if newHead.value != 5 {
		panic("WA")
	}
	if newHead.right != &node5 {
		panic("WA")
	}
	if newHead.right.value != 8 {
		panic("WA")
	}
}