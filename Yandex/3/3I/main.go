package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solution(n int, arr []int) string {
	buf := make([]ks, 10_000)
	for i := 0; i < len(arr); i++ {
		a := arr[i]
		idxArr := len(arr)
		if a > 0 {
			idxArr = i
		}
		k := ks{
			k:      a,
			idxArr: idxArr,
			n:      buf[a].n + 1,
		}
		buf[a] = k
	}

	sort.Slice(buf, func(i, j int) bool {
		if buf[i].n > buf[j].n {
			return buf[i].n > buf[j].n
		}
		if buf[i].n == buf[j].n {
			return buf[i].idxArr < buf[j].idxArr
		}
		return false
	})

	var ans []string
	for i := 0; i < len(buf); i++ {
		ans = append(ans, strconv.Itoa(buf[i].k))
		if len(ans) == n {
			break
		}
	}

	return strings.Join(ans, " ")
}

type ks struct {
	k      int
	n      int
	idxArr int
}

func main() {
	scanner := makeScanner()
	readInt(scanner)
	arr := readArray(scanner)
	n := readInt(scanner)
	printArray(solution(n, arr))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 64 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func printArray(ans string) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(ans)

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
