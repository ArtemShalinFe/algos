package main

import (
	"bufio"
	"os"
	"strconv"
)

func solution(arr []string) []string {
	var ans []string
	m := make(map[string]string, 0)

	for i := 0; i < len(arr); i++ {
		if _, ok := m[arr[i]]; ok {
			continue
		}
		m[arr[i]] = arr[i]
		ans = append(ans, arr[i])
	}

	return ans
}

func main() {
	scanner := makeScanner()
	n := readInt(scanner)
	arr := readArray(n, scanner)
	printArray(solution(arr))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func printArray(m []string) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(m); i++ {
		writer.WriteString(m[i])
		writer.WriteString("\n")
	}
	writer.Flush()
}

func readArray(n int, scanner *bufio.Scanner) []string {
	arr := make([]string, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		arr[i] = scanner.Text()
	}
	return arr
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}
