package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solution(m int, edges [][]int) [][]int {
	colors := make([]string, m+1)
	vertices := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		colors[i] = "white"
		vertices[i] = make([]int, 0)
	}

	for i := 0; i < len(edges); i++ {
		e := edges[i]
		vertices[e[0]] = append(vertices[e[0]], e[1])
		vertices[e[1]] = append(vertices[e[1]], e[0])
	}

	var stack []int
	var visited [][]int

	for i := 1; i <= m; i++ {
		stack = append(stack, i)
		g := make([]int, 0)
		for len(stack) > 0 {
			v := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			color := colors[v]

			if color == "white" {
				colors[v] = "grey"
				stack = append(stack, v)

				g = append(g, v)

				for i := 0; i < len(vertices[v]); i++ {
					if colors[vertices[v][i]] == "white" {
						stack = append(stack, vertices[v][i])
					}
				}
			} else if color == "grey" {
				colors[v] = "black"
			} else {
				continue
			}
		}

		sort.Slice(g, func(i, j int) bool {
			return g[i] < g[j]
		})
		if len(g) > 0 {
			visited = append(visited, g)
		}
	}

	return visited
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
	writer.WriteString(strconv.Itoa(len(arr)))
	writer.WriteString("\n")

	for i := 0; i < len(arr); i++ {
		if len(arr[i]) == 0 {
			continue
		}

		for j := 0; j < len(arr[i]); j++ {
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
