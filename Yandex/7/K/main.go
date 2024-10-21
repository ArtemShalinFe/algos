package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func solution(a []int, b []int) (int, []int, []int) {
	ans := 0
	ansA := make([]int, 0)
	ansB := make([]int, 0)

	dp := make([][]int, 0)

	for i := 0; i < len(a)+1; i++ {
		dp = append(dp, make([]int, len(b)+1))
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			iIdx := i + 1
			jIdx := j + 1
			if a[i] == b[j] {
				dp[iIdx][jIdx] = dp[iIdx-1][jIdx-1] + 1
			} else {
				dp[iIdx][jIdx] = max(dp[iIdx][jIdx-1], dp[iIdx-1][jIdx])
			}
		}
	}
	ans = dp[len(a)][len(b)]

	i := len(a)
	j := len(b)
	for dp[i][j] != 0 {
		if a[i-1] == b[j-1] {
			ansA = append(ansA, i)
			ansB = append(ansB, j)
			i--
			j--
		} else {
			if dp[i][j] == dp[i-1][j] {
				i--
			} else if dp[i][j] == dp[i][j-1] {
				j--
			}
		}
	}

	return ans, ansA, ansB
}

func main() {
	scanner := makeScanner()
	_, a := readArray(scanner)
	_, b := readArray(scanner)
	print(solution(a, b))
}

func readArray(scanner *bufio.Scanner) (int, []int) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	arr := make([]int, n)

	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	for j := 0; j < len(listString); j++ {
		arr[j], _ = strconv.Atoi(listString[j])
	}
	return n, arr
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func print(count int, a []int, b []int) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(strconv.Itoa(count))
	writer.WriteString("\n")
	for i := len(a) - 1; i >= 0; i-- {
		writer.WriteString((strconv.Itoa(a[i])))
		writer.WriteString(" ")
	}
	if len(a) > 0 {
		writer.WriteString("\n")
	}

	for i := len(b) - 1; i >= 0; i-- {
		writer.WriteString((strconv.Itoa(b[i])))
		writer.WriteString(" ")
	}
	if len(b) > 0 {
		writer.WriteString("\n")
	}

	writer.Flush()
}
