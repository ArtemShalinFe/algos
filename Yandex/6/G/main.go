package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solution(m int, start int, edges [][]int) int {
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
		vertices[e[1]] = append(vertices[e[1]], e[0])
	}

	var planned []int
	var distance = make([]int, m+1)
	var ans int

	distance[start] = 0
	planned = append(planned, start)
	for len(planned) > 0 {
		u := planned[0]
		planned = planned[1:]

		vs := vertices[u]
		sort.Slice(vs, func(i, j int) bool {
			return vs[i] < vs[j]
		})

		for i := 0; i < len(vs); i++ {
			v := vertices[u][i]
			if colors[v] == "white" {
				distance[v] = distance[u] + 1
				colors[v] = "gray"
				planned = append(planned, v)
			}
		}
		colors[u] = "black"
	}

	for i := 0; i < len(distance); i++ {
		ans = max(ans, distance[i])
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

func printArray(ans int) {
	writer := bufio.NewWriter(os.Stdout)

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteString(" ")

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

	scanner.Scan()
	start, _ := strconv.Atoi(scanner.Text())

	return m, start, arr
}
