package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solution(m int, start int, edges [][]int) [][]int {
	colors := make([]string, m+1)
	for i := 1; i < len(colors); i++ {
		colors[i] = "white"
	}

	vertices := make([][]int, m+1)
	for i := 0; i < len(vertices); i++ {
		vertices[i] = make([]int, 0)
	}
	for i := 0; i < len(edges); i++ {
		e := edges[i]
		vertices[e[0]] = append(vertices[e[0]], e[1])
	}

	var stack []int
	entry := make([]int, m+1)
	leave := make([]int, m+1)
	time := 0

	stack = append(stack, start)
	for len(stack) > 0 {
		v := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		color := colors[v]

		if color == "white" {
			entry[v] = time
			colors[v] = "grey"
			stack = append(stack, v)

			vs := vertices[v]
			sort.Slice(vs, func(i, j int) bool {
				return vs[i] > vs[j]
			})

			for i := 0; i < len(vertices[v]); i++ {
				if colors[vertices[v][i]] == "white" {
					stack = append(stack, vertices[v][i])
				}
			}
			time++
		} else if color == "grey" {
			leave[v] = time
			colors[v] = "black"
			time++
		} else {
			continue
		}
	}

	ans := make([][]int, m)
	for j := 0; j < m; j++ {
		i := j + 1
		ans[j] = make([]int, 2)
		ans[j][0] = entry[i]
		ans[j][1] = leave[i]
	}

	return ans
}

func main() {
	scanner := makeScanner()
	m, start, arr := readArray(scanner)
	printArray(solution(m, start, arr))
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

func readArray(scanner *bufio.Scanner) (int, int, [][]int) {
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

	return m, 1, arr
}
