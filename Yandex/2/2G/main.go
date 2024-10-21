package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func solution(cmds []string) []string {
	var ans []string
	s := NewStack()

	for i := 0; i < len(cmds); i++ {
		cmd := cmds[i]
		ans = s.do(cmd, ans)
	}
	return ans
}

func main() {
	scanner := makeScanner()
	cmds := readCmds(scanner)
	printArray(solution(cmds))
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

func readCmds(scanner *bufio.Scanner) []string {
	n := readInt(scanner)
	cmds := make([]string, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		cmds[i] = scanner.Text()
	}
	return cmds
}

const (
	None = "None"
	errs = "error"
)

type Stack struct {
	items []int
}

func NewStack() *Stack {
	return &Stack{
		items: make([]int, 0),
	}
}

func (s *Stack) do(cmd string, res []string) []string {
	switch {
	case strings.Contains(cmd, "get_max"):
		if s.size() == 0 {
			res = append(res, None)
		} else {
			res = append(res, strconv.Itoa(s.getMax()))
		}
	case strings.Contains(cmd, "push"):
		value := strings.Split(cmd, " ")[1]
		num, _ := strconv.Atoi(value)
		s.push(num)
	case strings.Contains(cmd, "pop"):
		if s.size() == 0 {
			res = append(res, errs)
		} else {
			s.pop()
		}
	default:
		println("unknow cmd")
	}

	return res
}

func (s *Stack) push(x int) {
	s.items = append(s.items, x)
}

func (s *Stack) pop() {
	s.items = s.items[:len(s.items)-1]
}

func (s *Stack) size() int {
	return len(s.items)
}

func (s *Stack) getMax() int {
	var m *int
	for i := 0; i < len(s.items); i++ {
		if m == nil {
			m = &s.items[i]
		} else {
			l := max(s.items[i], *m)
			m = &l
		}
	}
	return *m
}
