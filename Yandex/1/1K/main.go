package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Solution(ints []int, n int) []int {
	var carry int
	carry = n

	for i := len(ints) - 1; carry > 0; i-- {
		if i < 0 {
			var t []int
			t = append(t, 0)
			t = append(t, ints[:]...)
			ints = t
			i = 0
		}

		lnum := ints[i]
		num := lnum + carry
		if num > 9 {
			carry = num % 10
			ints[i] = carry
			carry = (num - carry) / 10
		} else {
			carry = 0
			ints[i] = num
		}
	}

	return ints
}

func main() {
	scanner := makeScanner()
	readInt(scanner)
	ints := readArray(scanner)
	i := readInt(scanner)
	printArray(Solution(ints, i))
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
