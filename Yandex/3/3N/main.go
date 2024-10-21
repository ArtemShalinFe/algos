package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solution(arr [][]int) [][]int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})
	smllst := arr[0]

	var ans [][]int
	for i := 0; i < len(arr); i++ {
		if smllst[1] < arr[i][0] {
			ans = append(ans, smllst)
			smllst = arr[i]
		} else {
			if smllst[1] < arr[i][1] {
				smllst[1] = arr[i][1]
			}
		}
	}

	ans = append(ans, smllst)

	return ans
}

func main() {
	scanner := makeScanner()
	c := readInt(scanner)
	arr := readArray(c, scanner)
	printArray(solution(arr))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 64 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func printArray(ans [][]int) {
	writer := bufio.NewWriter(os.Stdout)

	for i := 0; i < len(ans); i++ {
		for j := 0; j < len(ans[i]); j++ {
			writer.WriteString(strconv.Itoa(ans[i][j]))
			writer.WriteString(" ")
		}
		writer.WriteString("\n")
	}

	writer.Flush()
}

func readArray(n int, scanner *bufio.Scanner) [][]int {
	var arr [][]int

	for i := 0; i < n; i++ {
		scanner.Scan()
		listString := strings.Split(scanner.Text(), " ")

		a := make([]int, 2)
		a[0], _ = strconv.Atoi(listString[0])
		a[1], _ = strconv.Atoi(listString[1])

		arr = append(arr, a)
	}

	return arr
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}
