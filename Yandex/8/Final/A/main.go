/*
https://contest.yandex.ru/contest/26133/run-report/117154439/

### Принцип работы

	Программа определяет общий префикс распакованных строк с помощью функций распаковки строки и функции определния
	общего префикса строк.

### Доказательство корректности

	Распаковка происходит с помощью рекурсивного подхода, при котором программа анализирует символы строки и
	- если символ это число, то его нужно запомнить в переменную repeat
	- если это любой другой символ, кроме скобок '[',']' и цифр, то он помещается в массив префиксов
	- если открывающая скобка '[', то запускается шаг цикл рекурсии, а после выхода из него результат добавляет к префиксам,
	  дублируется repeat раз
	- если это закрывающая скобка, то добавляет префиксы в результат, если они есть, и возвращает результат

	После распаковки программа определяет общий префикс. Для этого она в качестве базовой строки берет первую строку,
	каждую следующую строку сравнивает с базовой и ищет одинаковые символы с начала строк. В случае если программа нашла
	символы, значения которых не совпадают, она обрезает базовый элемент. Эти действия повторяются для каждой входной
	строки и в результате программа получает общий префикс всех строк.

### Временная сложность

	Распаковка строк: O(N * L), где N - количество распаковываемых строк, L - самая длинная строка.
	Поиск общего префикса строк происходит за O(N * L), где N - количество строк, а L - самая длинная строка.

	Временная сложность: O(N * L)

### Пространственная сложность

	Распаковка строки занимает O(L*D) памяти, где L - самая длинная строка, D - глубина рекурсии
	Общий префикс в худшем случае занимает O(K) памяти, где K - количество символов общего префикса

	Пространственная сложность: O(L*D + K)
*/
package main

import (
	"bufio"
	"os"
	"strconv"
)

func makeScanner() *bufio.Scanner {
	const maxCapacity = 5 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func main() {
	scanner := makeScanner()
	prefix := solution(scanner)
	printAns(prefix)
}

func solution(scanner *bufio.Scanner) string {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	prefix := ""
	if n > 0 {
		scanner.Scan()
		prefix = unpackString(scanner.Text())
	}

	for i := 1; i < n; i++ {
		scanner.Scan()
		s := unpackString(scanner.Text())
		prefix = s[:findNotEqualIdx(prefix, s)]
	}

	return prefix
}

func findNotEqualIdx(s1 string, s2 string) int {
	var i int
	for i < max(len(s1), len(s2)) {
		if i >= len(s1) || i >= len(s2) || s1[i] != s2[i] {
			break
		}

		i++
	}

	return i
}

func unpackString(s string) string {
	runes, _ := unpack([]rune(s), 0)
	unpackedString := ""
	for _, r := range runes {
		unpackedString = unpackedString + string(r)
	}

	return unpackedString
}

func unpack(s []rune, idx int) ([][]rune, int) {
	if s[idx] == '[' {
		idx++
	}

	repeat := 1
	var prefix []rune
	var result [][]rune
	for ; idx < len(s); idx++ {
		switch {
		case s[idx] == '[':
			runes, lastIdx := unpack(s, idx)

			if prefix != nil {
				result = append(result, prefix)
			}
			for i := 0; i < repeat; i++ {
				result = append(result, runes...)
			}

			idx = lastIdx
			prefix = nil
			repeat = 1
		case s[idx] == ']':
			if prefix != nil {
				result = append(result, prefix)
			}
			return result, idx
		default:
			if isDigit(string(s[idx])) {
				repeat, _ = strconv.Atoi(string(s[idx]))
			} else {
				prefix = append(prefix, s[idx])
			}
		}
	}
	if prefix != nil {
		result = append(result, prefix)
	}
	return result, idx
}

func isDigit(str string) bool {
	if _, err := strconv.Atoi(str); err != nil {
		return false
	}
	return true
}

func printAns(s string) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(s)
	writer.Flush()
}
