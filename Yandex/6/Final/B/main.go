/*
https://contest.yandex.ru/contest/25070/run-report/116479316/

### Принцип работы

	Алгоритм представляет карту железных дорог в виду графа, у которого вершины - это города, а ребра - дороги между городами.
	Причем, дороги типа R - будут являться прямыми дорогами от одного города к другому, а дороги типа B - будут являться
	обратными дорогами между этими же городами.

	Идея заключается в том, чтобы использовать алгоритм поиска в глубину (DFS) для обхода графа и проверять, что между
	парами вершин нет циклов.

### Доказательство корректности

	В качестве механизма определения посещенных вершин, алгоритм использует цвета:
	- white вершина еще не посещена, является значением по умолчанию;
	- gray - вершина посещена, но обработаны не все ее ребра, устанавливается этот цвет когда алгоритм впервые посещает вершину;
	- black - вершина посещена и все ее ребра обработаны.

	Алгоритм обходит граф и, если во время обхода графа он встречает вершину цвета `gray`, это означает, что вершина была
	посещена ранее в текущем обходе и в графе есть цикл. А цикл в графе в свою очередь говоритм алгоритму о том, что
	между вершинами графа существует несколько путей разных типов, а значит предлагаемая карта неоптимальная.

### Временная сложность

	Временная сложность складывается из:
	- Построения матрицы смежности состоящей из городов (V) и дорог (E), которая в сумме составляет O(V*E).
	- Обход графа в глубину О(V+E), где города (V), а дороги (E)

	Итого:
	- Временная сложность - O(V*E)+О(V+E)

### Пространственная сложность

	В памяти программы хранится:
	- Матрица смежности, размера V^2, где V - количество городов.
	- Список посещенных городов, в худшем случае там окажутся все города V.
	- При обходе графа программа хранит в памяти стек городов, в худшем случае там окажутся все города V.

	Итого пространственная сложность:
	- O(V^2 + V)
*/
package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func solution(n int, mor *MatrixOfRailways) bool {
	for v := 1; v <= n; v++ {
		if mor.colors[v] != "white" {
			continue
		}

		stack := []int{v}
		for len(stack) > 0 {
			peak := stack[len(stack)-1]

			if mor.colors[peak] == "white" {
				mor.colors[peak] = "gray"

				for i := 0; i < len(mor.matrix[peak]); i++ {
					u := mor.matrix[peak][i]
					switch mor.colors[u] {
					case "white":
						stack = append(stack, u)
					case "gray":
						return false
					}
				}
			} else {
				mor.colors[peak] = "black"
				stack = stack[:len(stack)-1]
			}
		}
	}

	return true
}

func main() {
	scanner := makeScanner()
	n := readInt(scanner)
	mor := readArray(n, scanner)
	printRes(solution(n, mor))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func printRes(res bool) {
	writer := bufio.NewWriter(os.Stdout)
	if res {
		writer.WriteString("YES")
	} else {
		writer.WriteString("NO")
	}
	writer.Flush()
}

type MatrixOfRailways struct {
	matrix [][]int
	colors []string
}

func readArray(n int, scanner *bufio.Scanner) *MatrixOfRailways {
	mor := &MatrixOfRailways{
		matrix: make([][]int, n+1),
		colors: make([]string, n+1),
	}

	for v := 1; v < n; v++ {
		scanner.Scan()

		list := strings.Split(scanner.Text(), "")
		for e := 0; e < len(list); e++ {
			target := v + e + 1
			if list[e] == "R" {
				mor.matrix[v] = append(mor.matrix[v], target)
				mor.colors[v] = "white"
			} else {
				mor.matrix[target] = append(mor.matrix[target], v)
				mor.colors[target] = "white"
			}
		}
	}

	return mor
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}
