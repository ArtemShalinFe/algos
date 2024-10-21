package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Solution(matrix [][]int, rows, cols int) [][]int {
	var ans [][]int
	for i := 0; i < cols; i++ {
		ans = append(ans, make([]int, rows))
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			ans[j][i] = matrix[i][j]
		}
	}

	return ans
}

func main() {
	scanner := makeScanner()
	rows := readInt(scanner)
	col := readInt(scanner)
	matrix := readMatrix(scanner, rows)
	printArray(Solution(matrix, rows, col))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func printArray(arr [][]int) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			writer.WriteString(strconv.Itoa(arr[i][j]))
			writer.WriteString(" ")
		}
		writer.WriteString("\n")
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
