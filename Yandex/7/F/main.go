package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func solution(n int, k int) int {
	dp := make([]int, n)
	dp[0] = 1

	for i := 1; i < n; i++ {
		for j := max(0, i-k); j < i; j++ {
			dp[i] = (dp[i] + dp[j]) % 1000000007
		}
	}

	return dp[len(dp)-1]
}

func main() {
	scanner := makeScanner()
	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	n, _ := strconv.Atoi(listString[0])
	k, _ := strconv.Atoi(listString[1])
	print(solution(n, k))
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
