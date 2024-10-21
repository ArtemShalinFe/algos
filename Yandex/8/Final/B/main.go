/*
https://contest.yandex.ru/contest/26133/run-report/117156771/

### Принцип работы

	Программа определяет возможно ли разбить текст на слова из словоря используя префиксное дерево trie, с помощью
	которого проверяется соответствие шаблона.

### Доказательство корректности

	Программа строит префиксное дерево по данным из словаря. Для каждой последней буквы слова проставляется отметка,
	что узел терминальный (terminal = true), ребра префиксного дерева хранятся в хеш-таблице.

	Для проверки разбиения текста на отдельные слова из набора, программа использует метод динамического программирования,
	для этого объявляется массив dp размером исходного текста + 1, где будут отметки о том символ исходного текста является
	последним символом слова из словаря.

	Программа в цикле проходит по всем буквам исходного текста и ищет совпадения в trie, заполняя массив dp, если узел
	терминальный и всю предыдущую строку можно составить из набора слов в словаре.

	Если последний узел в массиве dp окажется терминальным, то текст возможно разбить на слова из словаря.

### Временная сложность

	Построение префиксного дерева требует O(N*L) ресурсов, где N - количество слов, а L - максимальная длина слова.
	Проверка, что текст можно разбить на слова, которая выполняет поиск узлов в префиксном дереве происходит за O(T*L),
	где T - длина базовой строки, L - максимальная длина слова .

	Временная сложность: O(N*L+T*L)

### Пространственная сложность

	В памяти программы хранится префиксное дерево размера O(N*L), где N - количество слов, а L - среднее количество символов в словах.
	Массив dp для динамического программирования требует дополнительное пространство O(T), где Т - длина базовой строки
	для проверки разбиения на подстроки.

	Пространственная сложность: O(N*L+T)
*/
package main

import (
	"bufio"
	"os"
	"strconv"
)

func makeRootNode() *node {
	return &node{
		terminal: false,
		prefixes: make(map[string]*node, 0),
	}
}

type node struct {
	terminal bool
	prefixes map[string]*node
}

func addNode(root *node, s string) {
	for i := 0; i < len(s); i++ {
		if _, ok := root.prefixes[string(s[i])]; !ok {
			root.prefixes[string(s[i])] = &node{
				terminal: false,
				prefixes: make(map[string]*node, 0),
			}
		}

		root = root.prefixes[string(s[i])]
	}

	root.terminal = true
}

func solution(str string, words []string) bool {
	root := makeRootNode()
	for i := 0; i < len(words); i++ {
		addNode(root, words[i])
	}

	dp := make([]bool, len(str)+1)
	dp[0] = true

	for i := 0; i < len(str); i++ {
		if !dp[i] {
			continue
		}

		current := root
		for j := i; j < len(str); j++ {
			sym := string(str[j])
			n, ok := current.prefixes[sym]
			if !ok {
				break
			}

			if n.terminal {
				dp[j+1] = true
			}

			current = n
		}
	}

	return dp[len(str)]
}

func main() {
	scanner := makeScanner()
	str := readString(scanner)
	words := readWords(scanner)
	ans := solution(str, words)
	writer := bufio.NewWriter(os.Stdout)
	if ans {
		writer.WriteString("YES")
	} else {
		writer.WriteString("NO")
	}
	writer.Flush()
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readString(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func readWords(scanner *bufio.Scanner) []string {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	var words []string
	for i := 0; i < n; i++ {
		words = append(words, readString(scanner))
	}

	return words
}
