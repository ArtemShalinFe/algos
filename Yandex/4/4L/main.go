package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func solution(n int, k int, str string) []int {
	var ans []int

	mp := make(map[int][]int)
	h := 0
	for i := 0; i+n < len(str); i++ {
		if i != 0 {

		}
		subStr := str[i : i+n]
		h = polyHash(10000000, 100000000003, subStr)

		if p, ok := mp[h]; ok {
			p = append(p, i)
			mp[h] = p
		} else {
			mp[h] = []int{i}
		}
		if len(mp[h]) == k {
			ans = append(ans, mp[h][0])
		}
	}
	return ans
}

func polyHash(a int, m int, s string) int {
	l := len(s)
	result := ((int(s[0]) * a) + int(s[1])) % m
	for i := 2; i < l; i++ {
		result = ((result * a) + int(s[i])) % m
	}
	return result
}

func main() {
	scanner := makeScanner()
	nk := readArray(scanner)
	str := readString(scanner)
	printRes(solution(nk[0], nk[1], str))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 128 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func printRes(m []int) {
	writer := bufio.NewWriter(os.Stdout)
	for j := 0; j < len(m); j++ {
		writer.WriteString(strconv.Itoa(m[j]))
		writer.WriteString(" ")
	}
	writer.Flush()
}

func readString(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
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
