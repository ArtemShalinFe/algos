package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func solution(d int, arr []int, n int) (int, int) {
	num1 := binarySearch(arr, n, 0, d)
	num2 := binarySearch(arr, n*2, 0, d)
	return num1, num2
}

func binarySearch(arr []int, n int, left int, right int) int {
	for left < right {
		mid := (left + right) / 2
		switch {
		case arr[mid] >= n && mid == 0:
			return mid + 1
		case arr[mid] >= n && n > arr[mid-1]:
			return mid + 1
		case arr[mid] >= n:
			right = mid
		default:
			left = mid + 1
		}
	}

	return -1
}

// func solutionRec(arr []int, n int) (int, int) {
// 	a := binarySearchRec(arr, n, 0, len(arr))
// 	b := binarySearchRec(arr, n*2, 0, len(arr))
// 	return a, b
// }

// func binarySearchRec(arr []int, n int, left int, right int) int {
// 	if right <= left {
// 		return -1
// 	}

// 	mid := (left + right) / 2
// 	if arr[mid] >= n && (n > arr[mid-1] || mid == 0) {
// 		return mid + 1
// 	} else if arr[mid] >= n {
// 		return binarySearchRec(arr, n, left, mid)
// 	} else {
// 		return binarySearchRec(arr, n, mid+1, right)
// 	}
// }

func main() {
	scanner := makeScanner()
	d := readInt(scanner)
	arr := readArray(scanner)
	p := readInt(scanner)
	printArray(solution(d, arr, p))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 64 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func printArray(a int, b int) {
	writer := bufio.NewWriter(os.Stdout)

	writer.WriteString(strconv.Itoa(a))
	writer.WriteString(" ")
	writer.WriteString(strconv.Itoa(b))
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
