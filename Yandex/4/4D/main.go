package main

import (
	"bufio"
	"os"
	"strconv"
)

func solution(a int, m int, s string) int {
	l := len(s)

	if l == 0 {
		return 0
	}

	if l == 1 {
		return int(s[0]) % m
	}

	result := ((int(s[0]) * a) + int(s[1])) % m
	for i := 2; i < l; i++ {
		result = ((result * a) + int(s[i])) % m
	}
	return result
}

func main() {
	scanner := makeScanner()
	a := readInt(scanner)
	m := readInt(scanner)
	s := readString(scanner)
	printRes(solution(a, m, s))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func printRes(m int) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(strconv.Itoa(m))
	writer.Flush()
}

func readString(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}
