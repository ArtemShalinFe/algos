// https://contest.yandex.ru/contest/22450/run-report/114060387/
package main

import (
	"bufio"
	"os"
	"strconv"
)

// Вычислительная сложность - квадратичная О(n)2 - из-за полного обхода field.
// Пространственная сложность - константная - О(1)
func solution(k int, field [][]int) int {
	var ans int

	arr := make([]int, 10)
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			arr[field[i][j]] += 1
		}
	}

	s := k * 2
	for t := 1; t <= 9; t++ {
		if arr[t] > 0 && arr[t] <= s {
			ans++
		}
	}

	return ans
}

func main() {
	scanner := makeScanner()
	k := readInt(scanner)
	field := readField(scanner)
	printAnswer(solution(k, field))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func printAnswer(ans int) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(strconv.Itoa(ans))
	writer.WriteString(" ")
	writer.Flush()
}

func readField(scanner *bufio.Scanner) [][]int {
	var field [][]int
	for i := 0; i <= 4; i++ {
		field = append(field, readArray(scanner))
	}

	return field
}

func readArray(scanner *bufio.Scanner) []int {
	scanner.Scan()
	text := scanner.Text()
	arr := make([]int, len(text))

	for i := 0; i < len(text); i++ {
		w := string(text[i])
		if w != "." {
			arr[i], _ = strconv.Atoi(w)
		}
	}
	return arr
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}
