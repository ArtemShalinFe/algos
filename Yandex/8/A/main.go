package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	scanner := makeScanner()
	strs := readArray(scanner)
	writer := bufio.NewWriter(os.Stdout)
	for i := len(strs) - 1; i >= 0; i-- {
		writer.WriteString(strs[i])
		writer.WriteString("\n")
	}
	writer.Flush()
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readArray(scanner *bufio.Scanner) []string {
	scanner.Scan()
	return strings.Split(scanner.Text(), " ")
}
