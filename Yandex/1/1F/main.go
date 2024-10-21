package main

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

func Solution(text string) bool {
	text = strings.ToLower(text)
	text = strings.Join(regexp.MustCompile(`[a-zA-Z0-9]+`).FindAllString(text, -1), "")
	count := len(text) - 1
	for i := 0; i < count; i++ {
		left := text[i]
		right := text[count-i]
		if left != right {
			return false
		}
	}

	return true
}

func main() {
	scanner := makeScanner()
	text := readString(scanner)
	printRes(Solution(text))
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
