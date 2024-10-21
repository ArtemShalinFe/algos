package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func solution(arr []int) []int {
	countedValues := make([]int, 3)
	for _, value := range arr {
		countedValues[value]++
	}

	index := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < countedValues[i]; j++ {
			arr[index] = i
			index++
		}
	}
	return arr
}

func main() {
	scanner := makeScanner()
	n := readInt(scanner)
	arr := readArray(n, scanner)
	printArray(n, solution(arr))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 64 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func printArray(n int, arr []int) {
	writer := bufio.NewWriter(os.Stdout)

	for i := 0; i < n; i++ {
		writer.WriteString(strconv.Itoa(arr[i]))
		writer.WriteString(" ")
	}
	writer.Flush()
}

func readArray(n int, scanner *bufio.Scanner) []int {
	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	arr := make([]int, len(listString))
	for i := 0; i < n; i++ {
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
