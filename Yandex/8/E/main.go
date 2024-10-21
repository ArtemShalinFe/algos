package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

type ins struct {
	text string
	pos  int
}

func main() {
	scanner := makeScanner()
	strs := readArray(scanner)
	n := readInt(scanner)

	var inss []ins
	for i := 0; i < n; i++ {
		scanner.Scan()
		t := strings.Split(scanner.Text(), " ")
		text := t[0]
		pos, _ := strconv.Atoi(t[1])
		inss = append(inss, ins{text: text, pos: pos})
	}

	sort.Slice(inss, func(i, j int) bool {
		return inss[i].pos < inss[j].pos
	})

	var res []string
	prev := 0
	for i := 0; i < len(inss); i++ {
		ins := inss[i]
		p := ins.pos
		res = append(res, strs[prev:p], ins.text)
		prev = p
	}

	res = append(res, strs[prev:])

	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(strings.Join(res, ""))
	writer.Flush()
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readArray(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
