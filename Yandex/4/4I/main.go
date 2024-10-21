package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solution(str []string) [][]int {
	mp := make(map[string][]int)
	for i := 0; i < len(str); i++ {
		s := strings.Split(str[i], "")
		sort.Strings(s)
		h := strings.Join(s, "")

		if a, ok := mp[h]; ok {
			a = append(a, i)
			sort.SliceStable(a, func(i, j int) bool {
				return i < j
			})
			mp[h] = a
		} else {
			mp[h] = []int{i}
		}
	}

	var ans [][]int
	for _, v := range mp {
		ans = append(ans, v)
	}

	sort.SliceStable(ans, func(i, j int) bool {
		return ans[i][0] < ans[j][0]
	})

	return ans
}

func main() {
	scanner := makeScanner()
	readInt(scanner)
	b := readArray(scanner)
	printRes(solution(b))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 64 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func printRes(m [][]int) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			writer.WriteString(strconv.Itoa(m[i][j]))
			writer.WriteString(" ")
		}
		writer.WriteString("\n")
	}

	writer.Flush()
}

func readArray(scanner *bufio.Scanner) []string {
	scanner.Scan()
	return strings.Split(scanner.Text(), " ")
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}
