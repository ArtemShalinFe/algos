package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solution(money []int, homes []int) int {
	var ans int

	lenHomes := money[0]
	budget := money[1]

	println(money[0])
	println(money[1])

	sort.Slice(homes, func(i, j int) bool {
		return homes[i] < homes[j]
	})

	for i := 0; i < lenHomes; i++ {
		home := homes[i]
		if home > budget {
			break
		}
		budget = budget - home
		println(budget)
		println(ans)
		ans++
	}

	return ans
}

func main() {
	scanner := makeScanner()
	money := readArray(scanner)
	homes := readArray(scanner)
	printRes(solution(money, homes))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 64 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func printRes(n int) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(strconv.Itoa(n))
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
