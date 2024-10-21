package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Solution(a string, b string) string {
	var ans []string

	res := 0
	cry := 0

	lenA := len(a) - 1
	lenB := len(b) - 1

	for {
		if lenA < 0 && lenB < 0 {
			break
		}

		intA := 0
		intB := 0

		if lenA >= 0 {
			intA, _ = strconv.Atoi(string(a[lenA]))
			lenA--
		}

		if lenB >= 0 {
			intB, _ = strconv.Atoi(string(b[lenB]))
			lenB--
		}

		res, cry = binSum(intA, intB, cry)
		ans = append(ans, strconv.Itoa(res))
	}

	if cry > 0 {
		ans = append(ans, "1")
	}

	var result []string
	for i := len(ans) - 1; i >= 0; i-- {
		result = append(result, ans[i])
	}
	return strings.Join(result, "")
}

func binSum(a, b, c int) (int, int) {
	binSum := a + b + c

	switch {
	case binSum == 0:
		return 0, 0
	case binSum == 1:
		return 1, 0
	case binSum == 2:
		return 0, 1
	case binSum == 3:
		return 1, 1
	default:
		return 0, 0
	}
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
