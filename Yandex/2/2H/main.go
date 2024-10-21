package main

import (
	"bufio"
	"os"
)

func Solution(brackets string) bool {
	stack := make([]string, 0)
	for i := 0; i < len(brackets); i++ {
		bracket := string(brackets[i])
		pair := getBracket(bracket)

		if pair != "" {
			stack = append(stack, pair)
		} else {
			if len(stack) > 0 && bracket == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}

func getBracket(b string) string {
	switch b {
	case "{":
		return "}"
	case "(":
		return ")"
	case "[":
		return "]"
	default:
		return ""
	}
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func printRes(res bool) {
	writer := bufio.NewWriter(os.Stdout)
	t := "False"
	if res {
		t = "True"
	}
	writer.WriteString(t)
	writer.WriteString(" ")
	writer.Flush()
}

func main() {
	scanner := makeScanner()
	brackets := readString(scanner)
	printRes(Solution(brackets))
}

func readString(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}
