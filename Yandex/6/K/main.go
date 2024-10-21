package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func solution(n int, _ int, mx [][]int) [][]int {
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		p := paths{
			n:    n,
			mx:   mx,
			dist: make([]int, n),
			//previous: make([]int, m),
			visited: make([]bool, n),
		}

		p.dijkstra(i)
		for i := 0; i < len(p.dist); i++ {
			if p.dist[i] == math.MaxInt32 {
				p.dist[i] = -1
			}
		}
		ans[i] = p.dist
	}
	return ans
}

type paths struct {
	n    int
	mx   [][]int
	dist []int
	//previous []int
	visited []bool
}

func (p *paths) relax(u, v int) {
	// Проверяем, не получился ли путь короче найденного ранее.
	if p.dist[v] > p.dist[u]+p.mx[u][v] {
		p.dist[v] = p.dist[u] + p.mx[u][v]
		//p.previous[v] = u
	}
}

func (p *paths) get_min_dist_not_visited_vertex() int {
	// Находим ещё непосещённую вершину с минимальным расстоянием от s.
	currentMinimum := math.MaxInt32
	currentMinimumVertex := -1

	for v := range p.n {
		if !p.visited[v] && p.dist[v] < currentMinimum {
			currentMinimum = p.dist[v]
			currentMinimumVertex = v
		}
	}
	return currentMinimumVertex
}

func (p *paths) dijkstra(s int) {
	for v := range p.n {
		p.dist[v] = math.MaxInt32 // Задаём расстояние по умолчанию.
		//p.previous[v] = -1        // Задаём предшественника для восстановления SPT.
		p.visited[v] = false // Список статусов посещённости вершин.
	}

	p.dist[s] = 0 // Расстояние от вершины до самой себя 0.

	for {
		u := p.get_min_dist_not_visited_vertex()

		if u == -1 || p.dist[u] == math.MaxInt32 {
			break
		}

		p.visited[u] = true
		// из множества рёбер графа выбираем те, которые исходят из вершины u
		neighbours := p.mx[u]
		for m, v := range neighbours {
			if v > 0 {
				p.relax(u, m)
			}
		}
	}
}

func main() {
	scanner := makeScanner()
	n, m, mx := readArray(scanner)
	printArray(solution(n, m, mx))
}

func readArray(scanner *bufio.Scanner) (int, int, [][]int) {
	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	n, _ := strconv.Atoi(listString[0])
	m, _ := strconv.Atoi(listString[1])

	mx := make([][]int, n)
	for i := 0; i < n; i++ {
		arr := make([]int, n)
		for k := 0; k < n; k++ {
			if k == i {
				arr[k] = 0
				continue
			}
			arr[k] = math.MaxInt32
		}
		mx[i] = arr
	}

	for j := 0; j < m; j++ {
		scanner.Scan()
		listString := strings.Split(scanner.Text(), " ")

		from, _ := strconv.Atoi(listString[0])
		to, _ := strconv.Atoi(listString[1])
		w, _ := strconv.Atoi(listString[2])

		from = from - 1
		to = to - 1

		mx[from][to] = min(w, mx[from][to])
		mx[to][from] = min(w, mx[to][from])
	}

	return n, m, mx
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
