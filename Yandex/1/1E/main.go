package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Word struct {
	Len   int
	Value string
}

const space = " "

func Solution(l int, text string) *Word {
	var savedWord []string
	var word []string
	text = strings.TrimSpace(text)
	for i := 0; i < len(text); i++ {
		s := string(text[i])
		if s == space {
			if len(word) > len(savedWord) {
				savedWord = word
			}
			word = make([]string, 0)
		} else {
			word = append(word, s)
		}
	}

	if len(word) > len(savedWord) {
		savedWord = word
	}

	return &Word{
		Len:   len(savedWord),
		Value: strings.Join(savedWord, ""),
	}
}

func main() {
	scanner := makeScanner()
	l := readInt(scanner)
	text := readString(scanner)
	printRes(Solution(l, text))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func printRes(res *Word) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(res.Value)
	writer.WriteString("\n")
	writer.WriteString(strconv.Itoa(res.Len))
	writer.WriteString(" ")
	writer.Flush()
}

func readString(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}
