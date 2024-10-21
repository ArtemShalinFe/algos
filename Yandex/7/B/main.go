package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solution(arr []*lesson) []*lesson {
	var ans []*lesson

	sort.Slice(arr, func(i, j int) bool {
		if arr[i].end == arr[j].end {
			return arr[i].start < arr[j].start
		}
		return arr[i].end < arr[j].end
	})

	var prev *lesson
	for i := 0; i < len(arr); i++ {
		cur := arr[i]
		if prev == nil {
			ans = append(ans, cur)
			prev = cur
			continue
		} else if prev.end <= cur.start {
			ans = append(ans, cur)
			prev = cur
		} else {
			continue
		}
	}

	return ans
}

func main() {
	scanner := makeScanner()
	ints := readArray(scanner)
	print(solution(ints))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func print(ans []*lesson) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(strconv.Itoa(len(ans)))
	writer.WriteString("\n")
	for i := 0; i < len(ans); i++ {
		l := ans[i]
		start := strconv.FormatFloat(l.start, 'f', -1, 64)
		writer.WriteString(start)
		writer.WriteString(" ")
		end := strconv.FormatFloat(l.end, 'f', -1, 64)
		writer.WriteString(end)
		writer.WriteString("\n")
	}

	writer.Flush()
}

func readArray(scanner *bufio.Scanner) []*lesson {
	scanner.Scan()
	count, _ := strconv.Atoi(scanner.Text())

	arr := make([]*lesson, count)
	for i := 0; i < count; i++ {
		scanner.Scan()
		listString := strings.Split(scanner.Text(), " ")
		start, _ := strconv.ParseFloat(listString[0], 32)
		end, _ := strconv.ParseFloat(listString[1], 32)
		arr[i] = &lesson{
			start: start,
			end:   end,
		}
	}
	return arr
}

type lesson struct {
	start float64
	end   float64
}
