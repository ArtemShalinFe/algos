/*
-- ССЫЛКА НА ОТЧЕТ --
https://contest.yandex.ru/contest/22781/run-report/114551449/

-- ПРИНЦИП РАБОТЫ --
Стек реализован на слайсе.
Если в процессе перебора входной последовательности алгоритм находит число, то добавляет его в конец стек,
если операцию над числами, то извлекает два последних элемента из стека, выполняет операцию над этими числами,
после чего результат помещает обратно в стек.

Когда обработана вся входная последовательность, алгоритм извлекает последний элемент из стека и возвращает его в качестве ответа.

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
Из описания алгоритма следует, что он реализует алгоритм вычисления значения выражения, записанного в обратной польской нотации.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
Время работы алгоритма зависит количества операций/операндов из входной строки размера n, поэтому временная сложность O(n).

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
Если входная последовательность содержит n чисел, то стек содержит n элементов.

Поэтому моя очередь будет потреблять O(n) памяти.
*/
package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := makeScanner()
	arr := readArray(scanner)
	printAns(solution(arr))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func printAns(ans int) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(strconv.Itoa(ans))
	writer.Flush()
}

func readArray(scanner *bufio.Scanner) []string {
	scanner.Scan()
	return strings.Split(scanner.Text(), " ")
}

const (
	add     = "+"
	subtr   = "-"
	multipl = "*"
	div     = "/"
)

func solution(arr []string) int {
	s := NewStack()

	for i := 0; i < len(arr); i++ {
		sym := string(arr[i])
		switch sym {
		case add:
			num1 := s.pop()
			num2 := s.pop()

			s.push(num2 + num1)
		case subtr:
			num1 := s.pop()
			num2 := s.pop()

			s.push(num2 - num1)
		case multipl:
			num1 := s.pop()
			num2 := s.pop()

			s.push(num2 * num1)
		case div:
			num1 := s.pop()
			num2 := s.pop()

			s.push(floorDiv(num2, num1))
		default:
			num, _ := strconv.Atoi(sym)
			s.push(num)
		}
	}
	return s.pop()
}

func floorDiv(num, divider int) int {
	if divider < 0 {
		num, divider = -num, -divider
	}

	remainder := num % divider
	if remainder < 0 {
		remainder += divider
	}

	return (num - remainder) / divider
}

type Stack struct {
	items []int
}

func NewStack() *Stack {
	return &Stack{
		items: make([]int, 0),
	}
}

func (s *Stack) push(v int) {
	s.items = append(s.items, v)
}

func (s *Stack) pop() int {
	tailIdx := len(s.items) - 1
	tailItem := s.items[tailIdx]
	s.items = s.items[:tailIdx]
	return tailItem
}
