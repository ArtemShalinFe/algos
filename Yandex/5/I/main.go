package main

import (
	"bufio"
	"os"
	"strconv"
)

func Solution(n int) int {
	var ans int

	if n == 0 {
		return 1
	}
	for i := 1; i <= n; i++ {
		ans = ans + Solution(i-1)*Solution(n-i)
	}

	return ans
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

func printRes(res int) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(strconv.Itoa(res))
	writer.Flush()
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}
