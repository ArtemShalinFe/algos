package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func solution(str []int, tmp []int) []int {
	ans := make([]int, 0)

	for i := 0; i < len(str)-len(tmp)+1; i++ {
		shift := str[i] - tmp[0]

		eq := true
		for j := 0; j < len(tmp); j++ {
			if str[i+j] != tmp[j]+shift {
				eq = false
				break
			}
		}
		if eq {
			ans = append(ans, i+1)
		}
	}

	return ans
}

func main() {
	scanner := makeScanner()
	str := readArray(scanner)
	tmp := readArray(scanner)
	ans := solution(str, tmp)

	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(ans); i++ {
		writer.WriteString(strconv.Itoa(ans[i]))
		writer.WriteString(" ")
	}
	writer.Flush()
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readArray(scanner *bufio.Scanner) []int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	sarr := strings.Split(scanner.Text(), " ")
	arr := make([]int, n)
	for i := 0; i < len(sarr); i++ {
		arr[i], _ = strconv.Atoi(sarr[i])
	}

	return arr
}
