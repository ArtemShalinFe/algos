package main

import (
	"bufio"
	"os"
	"strconv"
)

func Solution(n int) int {
	if n < 2 {
		return 1
	} else {
		return Solution(n-1) + Solution(n-2)
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

func printRes(n int) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(strconv.Itoa(n))
	writer.WriteString(" ")
	writer.Flush()
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}
