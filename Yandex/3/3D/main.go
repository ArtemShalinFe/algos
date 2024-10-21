package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solution(childCap []int, coockieSizes []int) int {
	var ans int

	sort.Slice(childCap, func(i, j int) bool {
		return childCap[i] < childCap[j]
	})

	sort.Slice(coockieSizes, func(i, j int) bool {
		return coockieSizes[i] < coockieSizes[j]
	})

	childIdx := 0
	cookieIdx := 0

forLoop:
	for childIdx < len(childCap) && cookieIdx < len(coockieSizes) {
		switch {
		case coockieSizes[cookieIdx] >= childCap[childIdx]:
			ans++
			cookieIdx++
			childIdx++
		case coockieSizes[cookieIdx] < childCap[childIdx]:
			cookieIdx++
		default:
			break forLoop
		}
	}

	return ans
}

func main() {
	scanner := makeScanner()
	childCount := readInt(scanner)
	childCap := readArray(childCount, scanner)
	coockieCount := readInt(scanner)
	coockieSizes := readArray(coockieCount, scanner)
	printRes(solution(childCap, coockieSizes))
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
