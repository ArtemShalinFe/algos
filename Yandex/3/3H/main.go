package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solution(arr []string) string {
	sort.Slice(arr, func(i, j int) bool {
		a, _ := strconv.Atoi(arr[i] + arr[j])
		b, _ := strconv.Atoi(arr[j] + arr[i])

		return a > b
	})
	return strings.Join(arr, "")
}

func main() {
	scanner := makeScanner()
	readInt(scanner)
	arr := readArray(scanner)
	printArray(solution(arr))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 64 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func printArray(ans string) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(ans)

	writer.Flush()
}

func readArray(scanner *bufio.Scanner) []string {
	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	return listString
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}
