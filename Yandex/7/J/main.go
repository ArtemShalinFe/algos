package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func solution(arr []int) []int {
	if len(arr) == 0 {
		return make([]int, 0)
	}

	dp := make([]int, len(arr))
	prev := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		dp[i] = 1
		prev[i] = -1
	}

	maxlen := 1
	lastIdx := 0

	for i := 1; i < len(arr); {
		for j := 0; j < i; {
			if arr[j] < arr[i] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
				prev[i] = j
			}
			j++
		}
		if dp[i] > maxlen {
			maxlen = dp[i]
			lastIdx = i
		}
		i++
	}

	var ans []int
	for lastIdx != -1 {
		ans = append(ans, lastIdx+1)
		lastIdx = prev[lastIdx]
	}
	return ans
}

func main() {
	scanner := makeScanner()
	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	n, _ := strconv.Atoi(listString[0])
	list := readArray(scanner, n)

	print(solution(list))
}

func readArray(scanner *bufio.Scanner, n int) []int {
	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	arr := make([]int, n)
	for j := 0; j < len(listString); j++ {
		v, _ := strconv.Atoi(listString[j])
		arr[j] = v
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

func print(ans []int) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(strconv.Itoa(len(ans)))
	writer.WriteString("\n")
	for i := len(ans) - 1; i >= 0; i-- {
		writer.WriteString(strconv.Itoa(ans[i]))
		writer.WriteString(" ")
	}

	writer.Flush()
}
