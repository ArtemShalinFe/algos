package main

import (
	"bufio"
	"os"
	"strconv"
)

// наивное решение, не прошло по времени
func Solution(n int) []int {
	var res []int

	for {
		if n == 1 {
			break
		}
		i := 2
		for {
			if n%i == 0 {
				res = append(res, i)
				n = n / i
				break
			}
			i++
		}
	}

	return res
}

func main() {
	scanner := makeScanner()
	n := readInt(scanner)
	printArray(Solution(n))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
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

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}
