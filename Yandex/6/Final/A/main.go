/*
https://contest.yandex.ru/contest/25070/run-report/116478406/

### Принцип работы

	Основная идея заключается в том, чтобы представить схему офиса в виде неориентированного графа. После чего найти вес
	максимального остовного дерева. При обходе графа алгоритм будет использовать такую структуру данных как приоритетная очередь,
	чтобы в качестве следующего ребра использовать ребро с максимальным весом.

### Доказательство корректности

	Алгоритм обходит граф до тех пор пока все связанные вершины не будут включены в максимальное остовного дерево.
	Во время обхода графа, благодаря приоритетной очереди, алгоритм всегда будет выбирать ребро с максимальным весом.
	В итоге алгоритм обойдет все связанные вершины по ребрам с максимальными весами и посчитает вес MST.

	Если после обхода графа имеются не посещенные вершины, то в графе есть несколько компонентов связности и программа выведет
	`Oops! I did it again`.

### Временная сложность

	Временная сложность состоит из:
	- Построения матрицы смежности O(V+E) из вершин (V) и ребер (E)
	- Нахождения максимального остовного дерева O(V+E) во время обхода графа программа добавляет и забирает ребро из
	приоритетной очереди за log E.

	Итого временная сложность:
	- O(V+E) + O(V+E*logE)

### Пространственная сложность

	В памяти программы хранится:
	- Матрица смежности, размера V^2, построенная из вершин V.
	- Массив посещенных вершин, в котором в худшем случае будут хранится все вершины V.
	- Приоритетная очередь, в которой в худшем случае будут хранится все вершины V.

	Итого пространственная сложность:
	- O(V^2 + V).
*/
package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func solution(mx [][]*Edge) string {
	var ans int
	if len(mx) == 0 {
		return "0"
	}
	pq := makeHeap()
	visited := make([]bool, len(mx)+1)
	visitedCount := 0

	visited[len(mx)-1] = true
	visitedCount++
	for i := 0; i < len(mx[len(mx)-1]); i++ {
		e := mx[len(mx)-1][i]
		if ok := visited[e.NodeID]; ok {
			continue
		}
		pq.Push(e)
	}

	for pq.Len() > 0 {
		maxEdge := pq.PopMax()
		if ok := visited[maxEdge.NodeID]; ok {
			continue
		}
		ans += maxEdge.Weight
		visited[maxEdge.NodeID] = true
		visitedCount++

		for i := 0; i < len(mx[maxEdge.NodeID]); i++ {
			e := mx[maxEdge.NodeID][i]
			if ok := visited[e.NodeID]; ok {
				continue
			}
			pq.Push(e)
		}
	}

	if visitedCount != len(mx)-1 {
		return "Oops! I did it again"
	}

	return strconv.Itoa(ans)
}

func readArray(scanner *bufio.Scanner) (int, int, [][]*Edge) {
	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")

	v, _ := strconv.Atoi(listString[0])
	e, _ := strconv.Atoi(listString[1])

	matrix := make([][]*Edge, v+1)
	for i := 1; i <= v; i++ {
		matrix[i] = make([]*Edge, 0)
	}

	for j := 1; j <= e; j++ {
		scanner.Scan()
		listString := strings.Split(scanner.Text(), " ")

		from, _ := strconv.Atoi(listString[0])
		to, _ := strconv.Atoi(listString[1])
		weight, _ := strconv.Atoi(listString[2])

		if from == to {
			continue
		}

		matrix[from] = append(matrix[from], &Edge{NodeID: to, Weight: weight})
		matrix[to] = append(matrix[to], &Edge{NodeID: from, Weight: weight})
	}

	return v, e, matrix
}

type Edge struct {
	NodeID int
	Weight int
}

func (e *Edge) isLower(j *Edge) bool {
	return e.Weight < j.Weight
}

type PriorityQueue struct {
	heap []*Edge
}

func makeHeap() *PriorityQueue {
	return &PriorityQueue{
		heap: make([]*Edge, 0),
	}
}

func (pq *PriorityQueue) Push(e *Edge) {
	pq.heap = append(pq.heap, e)
	pq.siftUp(pq.Len() - 1)
}

func (pq *PriorityQueue) Len() int {
	return len(pq.heap)
}

func (pq *PriorityQueue) PopMax() *Edge {
	result := pq.heap[0]
	pq.heap[0] = pq.heap[len(pq.heap)-1]
	pq.siftDown(0)
	pq.heap = pq.heap[:len(pq.heap)-1]

	return result
}

func (pq *PriorityQueue) siftDown(idx int) {
	for {
		left := 2 * idx
		right := 2*idx + 1

		if left >= len(pq.heap) {
			break
		}
		idxLrg := left
		if right < len(pq.heap) && pq.heap[left].isLower(pq.heap[right]) {
			idxLrg = right
		}

		if !pq.heap[idx].isLower(pq.heap[idxLrg]) {
			break
		}

		pq.heap[idx], pq.heap[idxLrg] = pq.heap[idxLrg], pq.heap[idx]
		idx = idxLrg
	}
}

func (pq *PriorityQueue) siftUp(idx int) int {
	if idx == 0 {
		return idx
	}

	for {
		parentIndex := idx / 2
		if !pq.heap[parentIndex].isLower(pq.heap[idx]) {
			break
		}
		pq.heap[parentIndex], pq.heap[idx] = pq.heap[idx], pq.heap[parentIndex]

		idx = parentIndex
	}

	return idx
}

func main() {
	scanner := makeScanner()
	n, m, list := readArray(scanner)
	if n > 1 && m == 0 {
		printAns("Oops! I did it again")
		return
	}
	printAns(solution(list))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func printAns(ans string) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(ans)
	writer.Flush()
}
