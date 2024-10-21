package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solution(borders []int) int {
	sort.Slice(borders, func(i, j int) bool {
		return borders[i] > borders[j]
	})
	var ans int

	for i := 2; i < len(borders); i++ {
		c := borders[i-2]
		b := borders[i-1]
		a := borders[i]
		if c < a+b && a+b+c > ans {
			ans = a + b + c
		}
	}

	return ans
}

func main() {
	scanner := makeScanner()
	readInt(scanner)
	borders := readArray(scanner)
	readInt(scanner)
	printRes(solution(borders))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 64 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func printRes(n int) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(strconv.Itoa(n))
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
