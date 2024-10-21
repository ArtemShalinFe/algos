package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var alph = "bdfhjlnprtvxz"

func solution(s []string, t []string) int {
	m := make(map[string]bool)
	for i := 0; i < len(alph); i++ {
		m[string(alph[i])] = true
	}

	sarr := make([]string, 0)
	for i := 0; i < len(s); i++ {
		if m[s[i]] {
			sarr = append(sarr, s[i])
		}
	}

	tarr := make([]string, 0)
	for i := 0; i < len(t); i++ {
		if m[t[i]] {
			tarr = append(tarr, t[i])
		}
	}

	ss := strings.Join(sarr, "")
	tt := strings.Join(tarr, "")

	if ss == tt {
		return 0
	}

	for i := 0; i < len(ss); i++ {
		if i == len(tt) {
			return 1
		}
		if ss[i] == tt[i] {
			continue
		}
		if ss[i] < tt[i] {
			return -1
		} else {
			return 1
		}
	}

	return -1
}

func main() {
	scanner := makeScanner()
	str1 := readString(scanner)
	str2 := readString(scanner)
	ans := solution(str1, str2)

	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(strconv.Itoa(ans))
	writer.Flush()
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readString(scanner *bufio.Scanner) []string {
	scanner.Scan()
	return strings.Split(scanner.Text(), "")
}
