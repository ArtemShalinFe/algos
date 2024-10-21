package main

import (
	"bufio"
	"os"
	"strconv"
)

func solution(s string) []int {
	n := len(s)
	pi := make([]int, n)
	pi[0] = 0
	for i := 1; i < n; i++ {
		k := pi[i-1]
		for k > 0 && s[k] != s[i] {
			k = pi[k-1]
		}
		if s[k] == s[i] {
			k++
		}
		pi[i] = k
	}
	return pi
}

func main() {
	scanner := makeScanner()
	str := readString(scanner)
	ans := solution(str)

	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(ans); i++ {
		writer.WriteString(strconv.Itoa(ans[i]))
		writer.WriteString(" ")
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

func readString(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}
