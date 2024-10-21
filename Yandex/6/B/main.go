package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func solution(m int, edges [][]int) [][]int {
	vertices := make([][]int, m+1)
	for i := 0; i < len(vertices); i++ {
		vertices[i] = make([]int, m+1)
	}
	for i := 0; i < len(edges); i++ {
		e := edges[i]
		vertices[e[0]][e[1]] = 1
	}
	return vertices
}

func main() {
	scanner := makeScanner()
	m, arr := readArray(scanner)
	printArray(solution(m, arr))
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
	for i := 1; i < len(arr); i++ {
		for j := 1; j < len(arr[i]); j++ {
			writer.WriteString(strconv.Itoa(arr[i][j]))
			writer.WriteString(" ")
		}
		writer.WriteString("\n")
	}
	writer.Flush()
}

func readArray(scanner *bufio.Scanner) (int, [][]int) {
	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	m, _ := strconv.Atoi(listString[0])
	n, _ := strconv.Atoi(listString[1])
	arr := make([][]int, n)

	for j := 0; j < n; j++ {
		scanner.Scan()
		listString := strings.Split(scanner.Text(), " ")

		from, _ := strconv.Atoi(listString[0])
		to, _ := strconv.Atoi(listString[1])

		es := arr[j]
		if es == nil {
			es = make([]int, 0)
		}
		es = append(es, from, to)

		arr[j] = es
	}

	return m, arr
}
