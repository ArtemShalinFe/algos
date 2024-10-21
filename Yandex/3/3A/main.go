package main

import (
	"bufio"
	"os"
	"strconv"
)

type App struct {
	res []string
}

/*
Если есть возможность поставить открывающую скобку, то мы ставим её.
Аналогично после этого если есть возможность поставить закрывающую скобку, то после этого мы ставим и её.
Таким образом строки будут выведены в лексографическом порядке,
так как сначала мы мы пытаемся поставить открывающую скобку.
При этом мы перебираем все возможные варианты последующих скобок для каждого возможного префикса 𝚊𝚗𝚜,
а следовательно в результате получаем все возможножные правильные скобочные последовательности
*/
func (a *App) solution(n int) []string {
	a.genBrackets(n, 0, 0, "")
	return a.res
}

func (a *App) genBrackets(n int, open int, close int, prefix string) {
	if open+close == 2*n {
		a.res = append(a.res, prefix)
		return
	}
	if open < n {
		a.genBrackets(n, open+1, close, prefix+"(")
	}
	if open > close {
		a.genBrackets(n, open, close+1, prefix+")")
	}
}

func main() {
	scanner := makeScanner()
	n := readInt(scanner)
	app := &App{
		res: make([]string, 0),
	}
	printArray(app.solution(n))
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
		writer.WriteString("\n")
	}
	writer.Flush()
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}
