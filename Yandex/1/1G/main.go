package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Solution(n int) string {
	var res []int
	for {
		res = append(res, n%2)
		n = n / 2
		if n == 0 {
			break
		}
	}

	var ans []string
	for i := len(res) - 1; i >= 0; i-- {
		ans = append(ans, strconv.Itoa(res[i]))
	}
	return strings.Join(ans, "")
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

func printRes(res string) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(res)
	writer.WriteString(" ")
	writer.Flush()
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}
