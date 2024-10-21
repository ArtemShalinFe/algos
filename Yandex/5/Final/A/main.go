/*
https://contest.yandex.ru/contest/24810/run-report/116073858/

### Принцип работы

	Алгоритм формирует из участников соревнований невозрастающую пирамиду. При получении участника из корня пирамиды
	алгоритм методом `heapify` восстанавливает правильный порядок элементов в пирамиде.

### Доказательство корректности

	При построении пирамиды из участников алгоритм строит дерево из всех элементов массива, не заботясь о соблюдении
	основного свойства кучи, а потом для всех вершин, у которых есть хотя бы один потомок вызывается метод `heapify`.
	Потомки гарантированно есть у первых heapSize/2 вершин.

	Метод `heapify` меняет местами вершину с наибольшим из потомков, пока не будет восстановлен правильный порядок элементов.
	В случае, если не найдется потомка, большего своего родителя процесс завершится. При сравнении участников метод использует
	функцию-компаратор, которая выше ставит того участника, у которого решено больше задач, в
	случае, если число решённых задач совпадает, первым будет идти участник с меньшим штрафом, в случае, если и штрафы
	совпадают, то первым будет тот, у которого логин идёт раньше в алфавитном (лексикографическом) порядке.

	При получении участника из пирамиды алгоритм всегда будет возвращать участника с корня пирамиды, так как это будет
	самый большой элемент из имеющихся.

### Временная сложность

	Временная сложность складывается из:
	- Построения невозрастающей пирамиды из массива участников соревнований размера O(N)
	- Получение самого большого элемента из невозрастающей пирамиды K*(Log K)

	Итого:
	- Временная сложность - O(N)+K*(Log K)

### Пространственная сложность

	В памяти программы хранится:
	- Массив, состоящий из участников соревнований N;
	- Невозрастающая пирамида, построенная из массива участников соревнований K.

	Размеры N и K одинаковые потому, что невозрастающая пирамида хранит данные в массиве, который был заполнен
	из массива участников соревнований. Так как K=N, то, пространственная сложность N+N=2N или просто N.

	Итого:
	- Пространственная сложность - O(N)
*/

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func makeHeap() *heap {
	return &heap{
		d: make([]*Participant, 0),
	}
}

type heap struct {
	d []*Participant
}

func (h *heap) heapify(idx int) {
	var leftChildIdx int
	var rightChildIdx int
	var largestChildIdx int
	heapSize := len(h.d)

	for {
		leftChildIdx = 2*idx + 1
		rightChildIdx = 2*idx + 2
		largestChildIdx = idx

		if leftChildIdx < heapSize && h.d[largestChildIdx].isLower(h.d[leftChildIdx]) {
			largestChildIdx = leftChildIdx
		}

		if rightChildIdx < heapSize && h.d[largestChildIdx].isLower(h.d[rightChildIdx]) {
			largestChildIdx = rightChildIdx
		}

		if largestChildIdx == idx {
			break
		}

		h.d[largestChildIdx], h.d[idx] = h.d[idx], h.d[largestChildIdx]
		idx = largestChildIdx
	}
}

func (h *heap) popMax() *Participant {
	// В ячейке 0 всегда самый большой элемент
	result := h.d[0]

	// Чтобы у дерева был корень - временно добавляем последний элемент в качестве корня дерева
	h.d[0] = h.d[len(h.d)-1]

	// "Тянем вниз по дереву" новый корень дерева, попутно поднимая вверх элементы, значения которых больше.
	h.heapify(0)

	// Элементов в куче стало меньше - нужно удалить последний элемент из дерева.
	h.d = h.d[:len(h.d)-1]

	return result
}

func solution(ps []*Participant) []*Participant {
	h := makeHeap()
	for i := 0; i < len(ps); i++ {
		h.d = append(h.d, ps[i])
	}
	for i := len(ps) / 2; i >= 0; i-- {
		h.heapify(i)
	}

	for i := 0; i < len(ps); i++ {
		ps[i] = h.popMax()
	}
	return ps
}

type Participant struct {
	name   string
	solved int
	fine   int
}

func (p *Participant) isLower(j *Participant) bool {
	if p.solved < j.solved {
		return true
	}

	if (p.solved == j.solved) && (p.fine > j.fine) {
		return true
	}

	if (p.solved == j.solved) && (p.fine == j.fine) {
		return p.name > j.name
	}

	return false
}

func main() {
	scanner := makeScanner()
	ps := readInput(scanner)
	printArray(solution(ps))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 64 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func printArray(ps []*Participant) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(ps); i++ {
		writer.WriteString(ps[i].name)
		writer.WriteString("\n")
	}
	writer.Flush()
}

func readInput(scanner *bufio.Scanner) []*Participant {
	scanner.Scan()
	stringInt := scanner.Text()
	n, _ := strconv.Atoi(stringInt)

	arr := make([]*Participant, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		listString := strings.Split(scanner.Text(), " ")

		solved, _ := strconv.Atoi(listString[1])
		fine, _ := strconv.Atoi(listString[2])

		arr[i] = &Participant{
			name:   listString[0],
			solved: solved,
			fine:   fine,
		}
	}
	return arr
}
