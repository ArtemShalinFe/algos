package main

import (
	"bufio"
	"os"
)

func Solution(a string, b string) string {
	var ans string

	mapa := make(map[string]int, 0)
	for i := 0; i < len(a); i++ {
		s := string(a[i])
		if s == " " {
			continue
		}
		mapa[s] = mapa[s] + 1
	}

	for i := 0; i < len(b); i++ {
		s := string(b[i])
		if s == " " {
			continue
		}

		if _, ok := mapa[s]; !ok {
			return s
		}
		mapa[s] = mapa[s] - 1
		if mapa[s] < 0 {
			return s
		}
	}

	return ans
}

func main() {
	scanner := makeScanner()
	a := readString(scanner)
	b := readString(scanner)
	printRes(Solution(a, b))
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
	writer.WriteString(" ")
	writer.Flush()
}

func readString(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}
