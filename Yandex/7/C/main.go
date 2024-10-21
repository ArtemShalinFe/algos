package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solution(m int, arr []*heap) int {
	var ans int

	sort.Slice(arr, func(i, j int) bool {
		if arr[i].sum == arr[j].sum {
			return arr[i].weight > arr[j].weight
		}
		return arr[i].sum > arr[j].sum
	})

	for i := 0; i < len(arr); i++ {
		h := arr[i]
		if (m - h.weight) >= 0 {
			ans = ans + (h.sum * h.weight)
			m = m - h.weight
		} else {
			ans = ans + m*h.sum
			break
		}
	}

	return ans
}

func main() {
	scanner := makeScanner()
	scanner.Scan()
	M, _ := strconv.Atoi(scanner.Text())
	ints := readArray(scanner)
	print(solution(M, ints))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func print(ans int) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(strconv.Itoa(ans))
	writer.WriteString("\n")

	writer.Flush()
}

func readArray(scanner *bufio.Scanner) []*heap {
	scanner.Scan()
	count, _ := strconv.Atoi(scanner.Text())

	arr := make([]*heap, count)
	for i := 0; i < count; i++ {
		scanner.Scan()
		listString := strings.Split(scanner.Text(), " ")
		sum, _ := strconv.Atoi(listString[0])
		weight, _ := strconv.Atoi(listString[1])
		arr[i] = &heap{
			weight: weight,
			sum:    sum,
		}
	}
	return arr
}

type heap struct {
	weight int
	sum    int
}
