package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Solution(matrix [][]int, rows, col, k, m int) []int {
	var res []int

	for i := max(0, k-1); i <= min(rows-1, k+1); i++ {
		if i != k {
			res = append(res, matrix[i][m])
		}
	}

	for j := max(0, m-1); j <= min(col-1, m+1); j++ {
		if j != m {
			res = append(res, matrix[k][j])
		}
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})

	return res
}

func main() {
	scanner := makeScanner()
	rows := readInt(scanner)
	col := readInt(scanner)
	matrix := readMatrix(scanner, rows)
	i := readInt(scanner)
	j := readInt(scanner)
	printArray(Solution(matrix, rows, col, i, j))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func printArray(arr []int) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(arr); i++ {
		writer.WriteString(strconv.Itoa(arr[i]))
		writer.WriteString(" ")
	}
	writer.Flush()
}

func readMatrix(scanner *bufio.Scanner, x int) [][]int {
	var matrix [][]int
	for i := 0; i < x; i++ {
		matrix = append(matrix, readArray(scanner))
	}

	return matrix
}

func readArray(scanner *bufio.Scanner) []int {
	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	arr := make([]int, len(listString))
	for i := 0; i < len(listString); i++ {
		arr[i], _ = strconv.Atoi(listString[i])
	}
	return arr
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}
