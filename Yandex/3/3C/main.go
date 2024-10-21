package main

import (
	"bufio"
	"os"
	"strings"
)

func solution(substr, str string) bool {
	if len(substr) > len(str) {
		return false
	}

	r := strings.Split(substr, "")
	ri := 0

	for i := 0; i < len(str); i++ {
		if string(str[i]) == string(r[ri]) {
			ri++
			if ri == len(substr) {
				return true
			}
		}
	}

	return ri == len(substr)
}

func main() {
	scanner := makeScanner()
	a := readString(scanner)
	b := readString(scanner)
	printRes(solution(a, b))
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

func readString(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}
