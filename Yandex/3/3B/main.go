package main

import (
	"bufio"
	"os"
)

type App struct {
	res []string
}

var m = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func (a *App) solution(str string) []string {
	a.gen(str, "")
	return a.res
}

func (a *App) gen(str string, comb string) {
	cmb := comb
	if len(comb) < len(str) {
		s := string(str[len(comb)])
		alph := m[s]
		for j := 0; j < len(alph); j++ {
			s := string(alph[j])
			cmb = comb + s
			a.gen(str, cmb)
		}
	} else {
		a.res = append(a.res, cmb)
	}
}

func main() {
	scanner := makeScanner()
	s := readString(scanner)
	app := &App{
		res: make([]string, 0),
	}
	printArray(app.solution(s))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func printArray(arr []string) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(arr); i++ {
		writer.WriteString(arr[i])
		writer.WriteString(" ")
	}
	writer.Flush()
}

func readString(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}
