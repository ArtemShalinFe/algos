package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func calculateY(a, b, c int64) string {
	if a%2 == 0 && b%2 == 0 && c%2 == 0 {
		return "WIN"
	} else if a%2 != 0 && b%2 != 0 && c%2 != 0 {
		return "WIN"
	} else {
		return "FAIL"
	}
}

func main() {
	scanner := makeScanner()
	ints := readArray(scanner)
	printRes(calculateY(ints[0], ints[1], ints[2]))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func printRes(res string) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(res)
	writer.Flush()
}

func readArray(scanner *bufio.Scanner) []int64 {
	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	arr := make([]int64, len(listString))
	for i := 0; i < len(listString); i++ {
		arr[i], _ = strconv.ParseInt(listString[i], 10, 64)
	}
	return arr
}
