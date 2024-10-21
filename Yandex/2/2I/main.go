package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func solution(size int, cmds []string) []string {
	var ans []string
	q := NewQueue(size)

	for i := 0; i < len(cmds); i++ {
		cmd := cmds[i]
		ans = q.do(cmd, ans)
	}
	return ans
}

func main() {
	scanner := makeScanner()
	cmds, size := readCmds(scanner)
	printArray(solution(size, cmds))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func printArray(arr []string) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(arr); i++ {
		writer.WriteString(arr[i])
		writer.WriteString("\n")
	}
	writer.Flush()
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}

func readCmds(scanner *bufio.Scanner) ([]string, int) {
	n := readInt(scanner)
	size := readInt(scanner)
	cmds := make([]string, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		cmds[i] = scanner.Text()
	}
	return cmds, size
}

const (
	None = "None"
	errs = "error"
)

type Queue struct {
	items   []int
	maxSize int
}

func NewQueue(maxSize int) *Queue {
	return &Queue{
		items:   make([]int, 0),
		maxSize: maxSize,
	}
}

func (q *Queue) do(cmd string, res []string) []string {
	switch {
	case strings.Contains(cmd, "size"):
		res = append(res, strconv.Itoa(q.size()))
	case strings.Contains(cmd, "push"):
		if q.size() == q.maxSize {
			res = append(res, errs)

		} else {
			value := strings.Split(cmd, " ")[1]
			num, _ := strconv.Atoi(value)
			q.push(num)
		}
	case strings.Contains(cmd, "pop"):
		if len(q.items) == 0 {
			res = append(res, None)
		} else {
			res = append(res, strconv.Itoa(q.pop()))
		}
	case strings.Contains(cmd, "peek"):
		if len(q.items) == 0 {
			res = append(res, None)
		} else {
			res = append(res, strconv.Itoa(q.peek()))
		}
	default:
		println("unknow cmd")
	}

	return res
}

func (q *Queue) push(x int) {
	q.items = append(q.items, x)
}

func (q *Queue) pop() int {
	n := q.items[0]
	q.items = q.items[1:]
	return n
}

func (q *Queue) peek() int {
	return q.items[0]
}

func (q *Queue) size() int {
	return len(q.items)
}
