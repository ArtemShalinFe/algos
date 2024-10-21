package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func solution(dp [][]int) (int, []string) {
	for i := len(dp) - 2; i >= 0; i-- {
		for j := 1; j < len(dp[i]); j++ {
			dp[i][j] = max(dp[i+1][j], dp[i][j-1]) + dp[i][j]
		}
	}

	ans := dp[0][len(dp[0])-1]
	var path []string

	for i, j := 0, len(dp[0])-1; i < len(dp)-1 && j > 0; {
		if dp[i][j-1] < dp[i+1][j] && i+1 < len(dp)-1 {
			path = append(path, "U")
			i++
			continue
		} else if dp[i][j-1] > dp[i+1][j] && j-1 > 0 {
			path = append(path, "R")
			j--
			continue
		} else if dp[i][j-1] == dp[i+1][j] {
			if i+1 < len(dp)-1 {
				path = append(path, "U")
				i++
				continue
			} else if j-1 > 0 {
				path = append(path, "R")
				j--
				continue
			} else {
				break
			}
		} else {
			break
		}
	}

	return ans, path
}

func main() {
	scanner := makeScanner()
	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	n, _ := strconv.Atoi(listString[0])
	m, _ := strconv.Atoi(listString[1])
	matrix := readArray(scanner, n, m)

	print(solution(matrix))
}

func readArray(scanner *bufio.Scanner, n int, m int) [][]int {
	arr := make([][]int, n+1)
	for i := 0; i < n; i++ {
		scanner.Scan()
		listString := strings.Split(scanner.Text(), "")
		arr2 := make([]int, m+1)
		for j := 0; j < len(listString); j++ {
			v, _ := strconv.Atoi(listString[j])
			arr2[j+1] = v
		}
		arr[i] = arr2
	}
	arr[n] = make([]int, m+1)
	return arr
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func print(ans int, path []string) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(strconv.Itoa(ans))
	writer.WriteString("\n")
	for i := len(path) - 1; i >= 0; i-- {
		writer.WriteString(path[i])
	}
	writer.Flush()
}
