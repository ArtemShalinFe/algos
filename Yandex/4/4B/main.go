package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func solution(arr []int) int {
	mCount := map[int]struct {
		first int
		last  int
	}{
		0: {0, 0},
	}
	count := 0
	ans := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			count++
		} else {
			count--
		}

		cur, ok := mCount[count]
		if !ok {
			mCount[count] = struct {
				first int
				last  int
			}{i + 1, i + 1}
		} else {
			cur.last = i + 1
			mCount[count] = cur
		}

		if (mCount[count].last - mCount[count].first) > ans {
			ans = mCount[count].last - mCount[count].first
		}
	}

	return ans
}

func main() {
	scanner := makeScanner()
	readInt(scanner)
	arr := readArray(scanner)
	printRes(solution(arr))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func printRes(m int) {
	writer := bufio.NewWriter(os.Stdout)
	// for i := 0; i < len(m); i++ {
	writer.WriteString(strconv.Itoa(m))
	// 	writer.WriteString("\n")
	// }
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

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}
