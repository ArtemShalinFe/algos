package main

import (
	"bufio"
	"os"
	"strings"
)

func solution(s []string, t []string) bool {
	var failures = 1
	var maxLen = max(len(s), len(t))
	i := 0
	j := 0

	for i < maxLen && j < maxLen {
		if s[i] == t[j] {
			i++
			j++
		} else {
			if len(s) == len(t) {
				i++
				j++
			} else {
				if len(s) > len(t) {
					i++
				} else {
					j++
				}
			}

			failures--
		}
	}

	return failures >= 0
}

func main() {
	scanner := makeScanner()
	str1 := readString(scanner)
	str2 := readString(scanner)
	ans := solution(str1, str2)
	writer := bufio.NewWriter(os.Stdout)
	if ans {
		writer.WriteString("OK")
	} else {
		writer.WriteString("FAIL")
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

func readString(scanner *bufio.Scanner) []string {
	scanner.Scan()
	return strings.Split(scanner.Text(), "")
}
