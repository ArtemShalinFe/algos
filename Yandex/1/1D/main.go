package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Solution(t int, arr []int) int {
	if len(arr) == 1 {
		return 1
	}

	var res int
	count := len(arr) - 1
	for cur := 0; cur < len(arr); cur++ {
		prev := cur - 1
		next := cur + 1

		switch {
		case prev >= 0 && next <= count:
			if (arr[cur] > arr[prev]) && (arr[cur] > arr[next]) {
				res++
			}
		case prev < 0 && next <= count:
			if arr[cur] > arr[next] {
				res++
			}
		case prev >= 0 && next > count:
			if arr[cur] > arr[prev] {
				res++
			}
		default:
			continue
		}
	}

	return res
}

func main() {
	scanner := makeScanner()
	t := readInt(scanner)
	arr := readArray(scanner)
	printRes(Solution(t, arr))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func printRes(res int) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(strconv.Itoa(res))
	writer.WriteString(" ")
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
