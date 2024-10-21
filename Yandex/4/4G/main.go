package main

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
)

func solution(s int, arr []int) [][]int {
	ans := make([][]int, 0)
	// mp := make(map[[4]int]struct{}, 0)

	sumMap := make(map[int]struct {
		i int
		j int
	}, 0)
	n := len(arr)
	slices.Sort(arr)
	for i := 0; i < n-3; i++ {
		for j := i + 1; j < n-2; j++ {
			sumMap[arr[i]+arr[j]] = struct {
				i int
				j int
			}{i: i, j: j}
		}
	}

	return ans
}

func quickSearch(arr []int, left int, right int, x int) int {
	mid := (left + right) / 2
	for left <= right {
		if arr[mid] == x {
			return mid
		}
		if arr[mid] > x {
			return quickSearch(arr, left, mid-1, x)
		} else {
			return quickSearch(arr, mid+1, right, x)
		}
	}

	return -1
}

func main() {
	scanner := makeScanner()
	n := readInt(scanner)
	s := readInt(scanner)
	arr := readArray(n, scanner)
	printRes(solution(s, arr))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func printRes(m [][]int) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(strconv.Itoa(len(m)))
	writer.WriteString("\n")
	for j := 0; j < len(m); j++ {
		for i := 0; i < len(m[j]); i++ {
			writer.WriteString(strconv.Itoa(m[j][i]))
			writer.WriteString(" ")
		}
		writer.WriteString("\n")
	}
	writer.Flush()
}

func readString(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}

func readArray(n int, scanner *bufio.Scanner) []int {
	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i], _ = strconv.Atoi(listString[i])
	}
	return arr
}
