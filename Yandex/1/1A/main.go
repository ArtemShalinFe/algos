package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func calculateY(a, x, b, c int) int {
	return a*(x*x) + (b * x) + c
}

func main() {
	scanner := makeScanner()
	ints := readArray(scanner)
	printY(calculateY(ints[0], ints[1], ints[2], ints[3]))
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
	listString := strings.Split(scanner.Text(), " ")
	arr := make([]int, len(listString))
	for i := 0; i < len(listString); i++ {
		arr[i], _ = strconv.Atoi(listString[i])
	}
	return arr
}
