package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var vertices [][]int
var order []int
var visited []bool

func solution(m int, edges [][]int) []int {
	order = make([]int, 0)
	vertices = make([][]int, 0)
	visited = make([]bool, m+1)

	for i := 0; i <= m+1; i++ {
		vertices = append(vertices, make([]int, 0))
	}
	for i := 0; i < len(edges); i++ {
		e := edges[i]
		vs := append(vertices[e[0]], e[1])
		vertices[e[0]] = vs
	}

	for i := m; i > 0; i-- {
		if !visited[i] {
			topSort(i)
		}
	}

	var ans []int
	for i := len(order); i > 0; i-- {
		ans = append(ans, order[i-1])
	}

	return ans
}

func topSort(v int) {
	outgoingEdges := vertices[v]
	for _, w := range outgoingEdges {
		if !visited[w] {
			topSort(w)
		}
	}
	visited[v] = true
	order = append(order, v)
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

func printArray(arr []int) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(arr); i++ {
		writer.WriteString(strconv.Itoa(arr[i]))
		writer.WriteString(" ")
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
