package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func solution(m int, weights []int) int {
	dp := make([]int, m+1)

	for i := 0; i < len(weights); i++ {
		for j := m; j >= weights[i]; j-- {
			dp[j] = max(dp[j], dp[j-weights[i]]+weights[i])
		}
	}

	return dp[m]
}

func main() {
	scanner := makeScanner()
	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	n, _ := strconv.Atoi(listString[0])
	m, _ := strconv.Atoi(listString[1])
	list := readArray(scanner, n)

	print(solution(m, list))
}

func readArray(scanner *bufio.Scanner, n int) []int {
	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	arr := make([]int, n+1)
	for j := 0; j < len(listString); j++ {
		v, _ := strconv.Atoi(listString[j])
		arr[j+1] = v
	}
	return arr
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func print(ans int) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(strconv.Itoa(ans))
	writer.WriteString("\n")

	writer.Flush()
}
