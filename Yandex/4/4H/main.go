package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func solution(str string) string {
	var maxLen int

	for i := 0; i < len(str); i++ {
		mp := make(map[string]struct{})
		mp[string(str[i])] = struct{}{}

		for j := i + 1; j < len(str); j++ {
			sym := string(str[j])
			if _, ok := mp[sym]; ok {
				break
			}
			mp[sym] = struct{}{}
		}
		maxLen = max(maxLen, len(mp))
	}

	return strconv.Itoa(maxLen)
}

func main() {
	scanner := makeScanner()
	b := readString(scanner)
	printRes(solution(b))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 64 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func printRes(m string) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(m)
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

func readString(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}
