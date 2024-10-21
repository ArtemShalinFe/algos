package main

import (
	"bufio"
	"os"
	"strconv"
)

func Solution(n int) bool {
	if n == 1 {
		return true
	}

	for {
		carry := n % 4
		if carry > 0 {
			return false
		}
		n = n / 4
		if n == 1 {
			return true
		}
	}
}

func main() {
	scanner := makeScanner()
	n := readInt(scanner)
	printRes(Solution(n))
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

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}
