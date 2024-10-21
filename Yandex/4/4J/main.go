package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// vector<int> LCIS(a: int[n], b: int[m])
//   for i = 1 to n
//     ind = 0 // позиция "лучшего" элемента в массиве b
//     best = 0 // значение динамики для "лучшего" элемента
//     for j = 1 to m
//       d[i][j] = d[i - 1][j] // НОВП на a[1..i - 1] и b[1..j] (без элемента a[i])
//       if a[i] == b[j] and d[i - 1][j] < best + 1 // используем a[i]-й элемент для увеличения НОВП
//         d[i][j] = best + 1
//         prev[j] = ind
//       if a[i] > b[j] and d[i - 1][j] > best // при следующем равенстве a[i] == b[j']
//         best = d[i - 1][j] // в best будет храниться "лучший" элемент
//         ind = j // b[ind] < b[j'] и d[i-1][ind] → max
//   // восстановление (по массиву b)
//   pos = 1 // ищем лучший элемент d[n][pos] → max
//   for j = 1 to m
//     if d[n][pos] < d[n][j]
//       pos = j
//   // проходим по массиву b, выписывая элементы НОВП
//   answer: vector<int>
//   while pos ≠ 0
//     answer.pushBack(b[pos])
//     pos = prev[pos]
//   return answer

func solution(a []int, b []int) int {
	d := make([][]int, 0)
	for i := 0; i < len(a); i++ {
		x := make([]int, len(b))
		for j := 0; j < len(b); j++ {
			x[j] = 0
		}
		d = append(d, x)
	}

	for i := 1; i < len(a); i++ {
		ind := 0
		best := 0
		for j := 1; j < len(b); j++ {
			d[i][j] = d[i-1][j]
			if a[i] == b[j] && d[i-1][j] < best+1 {
				d[i][j] = best + 1
			}
			if a[i] > b[j] && d[i-1][j] > best {
				best = d[i-1][j]
				ind = j
			}
		}
		print(ind)
	}

	return 0
}

func main() {
	scanner := makeScanner()
	readInt(scanner)
	arr1 := readArray(scanner)
	readInt(scanner)
	arr2 := readArray(scanner)
	printRes(solution(arr1, arr2))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func printRes(m int) {
	writer := bufio.NewWriter(os.Stdout)
	// for i := 0; i < len(m); i++ {
	writer.WriteString(strconv.Itoa(m))
	// 	writer.WriteString("\n")
	// }
	writer.Flush()
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
