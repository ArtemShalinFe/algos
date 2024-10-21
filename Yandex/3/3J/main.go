package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func solution(arr []int) [][]int {
	var ans [][]int

	for i := 0; i < len(arr); i++ {
		swpd := false
		for j := 0; j < len(arr)-1; j++ {
			prev := arr[j+1]
			cur := arr[j]
			if cur > prev {
				arr[j], arr[j+1] = prev, cur
				swpd = true
			}
		}
		if swpd {
			c := make([]int, len(arr))
			copy(c, arr)
			ans = append(ans, c)
		}
	}
	if len(ans) == 0 {
		ans = append(ans, arr)
	}

	return ans
}

func main() {
	scanner := makeScanner()
	readInt(scanner)
	arr := readArray(scanner)
	printArray(solution(arr))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 64 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func printArray(ints [][]int) {
	writer := bufio.NewWriter(os.Stdout)

	for i := 0; i < len(ints); i++ {
		for j := 0; j < len(ints[i]); j++ {
			writer.WriteString(strconv.Itoa(ints[i][j]))
			writer.WriteString(" ")
		}
		writer.WriteString("\n")
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
