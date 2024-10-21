package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func solution(a string, b string) string {
	if len(a) != len(b) {
		return "NO"
	}
	mAtoB := make(map[rune]rune)
	mBtoA := make(map[rune]rune)
	for i := 0; i < len(a); i++ {
		aSym := rune(a[i])
		bSym := rune(b[i])

		if v, ok := mAtoB[aSym]; ok {
			if v != bSym {
				return "NO"
			}
		} else {
			mAtoB[aSym] = bSym
		}

		if v, ok := mBtoA[bSym]; ok {
			if v != aSym {
				return "NO"
			}
		} else {
			mBtoA[bSym] = aSym
		}
	}

	return "YES"
}

func main() {
	scanner := makeScanner()
	a := readString(scanner)
	b := readString(scanner)
	printRes(solution(a, b))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
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
