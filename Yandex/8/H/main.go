package main

import (
	"bufio"
	"os"
)

func solution(str string, pattern string, ext string) string {
	var pi []int
	s := pattern + "#" + str
	π := make([]int, len(pattern))
	var πPrev int
	for i := 1; i < len(s); i++ {
		k := πPrev
		for k > 0 && s[k] != s[i] {
			k = π[k-1]
		}
		if s[k] == s[i] {
			k++
		}
		if i < len(pattern) {
			π[i] = k
		}
		πPrev = k
		if k == len(pattern) {
			pi = append(pi, i-2*len(pattern))
		}
	}

	shift := 0
	for i := 0; i < len(pi); i++ {
		tmp := str[:pi[i]+shift] + ext
		str = tmp + str[pi[i]+len(pattern)+shift:]
		shift = shift + len(ext) - len(pattern)
	}

	return str
}

func main() {
	scanner := makeScanner()

	str := readString(scanner)
	pattern := readString(scanner)
	ext := readString(scanner)

	ans := solution(str, pattern, ext)

	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(ans)
	writer.Flush()
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readString(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}
