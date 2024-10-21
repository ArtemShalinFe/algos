// https://contest.yandex.ru/contest/22450/run-report/114060179/
package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Вычислительная сложность - О(n) - В худшем случае когда в массиве только один 0 и он в конце, то есть О(2n), но 2 опускаем.
// Пространственная сложность - константная - О(1) - кажется, что дополнительная память выделяется только для переменной zeroPos.
func solution(arr []int) []int {
	inc := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			if inc < 0 {
				inc = i
			}
			for j := 1; j <= inc; j++ {
				k := i - j
				if arr[k] != -1 && arr[k] <= j {
					break
				}
				arr[k] = j
			}
			inc = 1
			continue
		}

		arr[i] = inc
		if inc > 0 {
			inc++
		}
	}

	return arr
}

func main() {
	scanner := makeScanner()
	readInt(scanner)
	arr := readArray(scanner)
	printArray(solution(arr))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 64 << (10 * 2) // 64Mb
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
