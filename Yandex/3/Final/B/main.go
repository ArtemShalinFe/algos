/*
https://contest.yandex.ru/contest/23815/run-report/115025226/

### Принцип работы

	Принцип работы заключен в рекурсивном вызове функции сортировки, которая применяется для левой и правой части массива.
	Рекурсивные вызовы происходят до тех пор, пока левая граница сортируемого массива меньше
	правой границы сортируемого массива.

	В качестве функции разделении массива используется схема Хоара:
	- В качестве опорного элемента выбирается элемент в середине сортируемой части.
	- Объявляются 2 указателя на левую и правую границу массива.
	- В цикле происходит просмотр с левой границы массива, который происходит до тех пор, пока не будет найден элемент,
	который больше опорного.
	- В цикле происходит просмотр с правой границы массива, который происходит до тех пор, пока не будет найден элемент,
	который меньше опорного.
	- После этого программа меняет найденные элементы местами.
	- В случае, если левая граница "перешла" правую функция возвращает указатель правой границы.

	Для сравнения элементов используется функция isLower, в которой реализована логика сравнения.

### Доказательство корректности

	Алгоритм является вариацией быстрой сортировки с разбиением Хоара и использованием функции-компаратора.

### Временная сложность

	При каждой итерации общая длина не отсортированного массива уменьшается на какую-то часть:
	- Временная сложность - O(N*log n) - логарифмическая
	Но в худшем случае, если элементы уже отсортированы или всегда выбирается не удачный опорный элемент:
	- Временная сложность - O(N^2) - квадратичная

### Пространственная сложность

	В памяти  программы хранится массив размера N, который был передан во входных данных.
	На каждой итерации цикла поочередно в рекурсивные вызовы копируется весь исходный массив,
	получается, что пространственная сложность зависит от размера исходного массива и глубины K рекурсивных вызовов.
	- Пространственная сложность - O(N*K)
*/
package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func solution(ps []Participant) []Participant {
	return quickSortHoar(ps, 0, len(ps)-1)
}

func quickSortHoar(ps []Participant, left int, right int) []Participant {
	if left < right {
		p := partitionHoar(ps, left, right)

		ps = quickSortHoar(ps, left, p)
		ps = quickSortHoar(ps, p+1, right)
	}

	return ps
}

func partitionHoar(ps []Participant, left int, right int) int {
	pivot := ps[(left+right)/2]
	i := left
	j := right

	for {
		for pivot.isLower(ps[i]) {
			i++
		}

		for ps[j].isLower(pivot) {
			j--
		}
		if i >= j {
			return j
		}
		ps[i], ps[j] = ps[j], ps[i]
	}
}

func (p Participant) isLower(j Participant) bool {
	if p.solved < j.solved {
		return true
	} else if (p.solved == j.solved) && (p.fine > j.fine) {
		return true
	} else if (p.solved == j.solved) && (p.fine == j.fine) {
		return p.name > j.name
	}

	return false
}

type Participant struct {
	name   string
	solved int
	fine   int
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

func printArray(ps []Participant) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(ps); i++ {
		writer.WriteString(ps[i].name)
		writer.WriteString("\n")
	}
	writer.Flush()
}

func readInput(scanner *bufio.Scanner) []Participant {
	scanner.Scan()
	stringInt := scanner.Text()
	n, _ := strconv.Atoi(stringInt)

	arr := make([]Participant, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		listString := strings.Split(scanner.Text(), " ")

		solved, _ := strconv.Atoi(listString[1])
		fine, _ := strconv.Atoi(listString[2])

		arr[i] = Participant{
			name:   listString[0],
			solved: solved,
			fine:   fine,
		}
	}
	return arr
}
