package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func solution(arr []int) int {
	var ans int
	for i := 1; i < len(arr); i++ {
		first := arr[i-1]
		second := arr[i]
		if first < second {
			ans = ans + (second - first)
		}
	}
	return ans
}

func main() {
	scanner := makeScanner()
	ints := readArray(scanner)
	printY(solution(ints))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func printY(y int) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(strconv.Itoa(y))
	writer.Flush()
}

func readArray(scanner *bufio.Scanner) []int {
	scanner.Scan()
	count, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	arr := make([]int, len(listString))
	for i := 0; i < count; i++ {
		arr[i], _ = strconv.Atoi(listString[i])
	}
	return arr
}
