package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func solution(a, m, n int, s string, pairs []pair) []int {
	if len(s) == 0 {
		return []int{0}
	}

	degree := degreeByModule(a, m, len(s))
	pref := prefixHash(s, a, m)
	result := make([]int, n)

	for i := 0; i < len(pairs); i++ {
		aI := pairs[i].left
		bI := pairs[i].right

		if aI-1 == 0 {
			result[i] = pref[bI]
		} else {
			v := pref[bI] - pref[aI-1]*degree[bI-aI+1]
			if v < 0 {
				v %= m
				v += m
				v %= m
			}
			result[i] = v
		}
	}

	return result
}

func degreeByModule(val, mod, n int) []int {
	degrees := make([]int, n+1)
	degrees[0] = 1

	for i := 1; i <= n; i++ {
		res := degrees[i-1] * val
		degrees[i] = res % mod
	}

	return degrees
}

func prefixHash(s string, a, m int) []int {
	pref := make([]int, len(s)+1)
	pref[0] = 0

	for i := 1; i <= len(s); i++ {
		pref[i] = (pref[i-1]*a + int(s[i-1])) % m
	}

	return pref
}

func main() {
	scanner := makeScanner()
	a := readInt(scanner)
	m := readInt(scanner)
	s := readString(scanner)
	n := readInt(scanner)
	pairs := readArray(n, scanner)
	printRes(solution(a, m, n, s, pairs))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func printRes(m []int) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(m); i++ {
		writer.WriteString(strconv.Itoa(m[i]))
		writer.WriteString("\n")
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

type pair struct {
	left  int
	right int
}

func readArray(n int, scanner *bufio.Scanner) []pair {
	arr := make([]pair, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		listString := strings.Split(scanner.Text(), " ")
		l, _ := strconv.Atoi(listString[0])
		r, _ := strconv.Atoi(listString[1])
		arr[i] = pair{
			left:  l,
			right: r,
		}
	}
	return arr
}
