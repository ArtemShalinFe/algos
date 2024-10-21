package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func (n *node) find(sym string) *node {
	for i := 0; i < len(n.nodes); i++ {
		root := n.nodes[i]
		if root.symbol == sym {
			return root
		}
	}

	return nil
}

func (n *node) add(sym string, word string) *node {
	for i := 0; i < len(n.nodes); i++ {
		root := n.nodes[i]
		if root.symbol == sym {
			root.words = append(root.words, word)
			return root
		}
	}

	nd := &node{
		symbol: sym,
		words:  []string{word},
		nodes:  make([]*node, 0),
	}
	n.nodes = append(n.nodes, nd)

	return nd
}

type node struct {
	symbol string
	words  []string
	nodes  []*node
}

func solution(data []string, reqs []string) []string {
	root := &node{
		symbol: "",
		words:  make([]string, 0),
		nodes:  make([]*node, 0),
	}

	for i := 0; i < len(data); i++ {
		nd := root
		word := string(data[i])
		for j := 0; j < len(word); j++ {
			sym := string(word[j])
			if sym == strings.ToUpper(sym) {
				nd = nd.add(sym, word)
			}
			root.words = append(root.words, word)
		}
	}

	m := make(map[string]string)
	var ans []string
	for i := 0; i < len(reqs); i++ {
		node := root
		req := reqs[i]
		syms := strings.Split(req, "")

		var wrds []string
		for i := 0; i < len(syms); i++ {
			node = node.find(syms[i])
			if node == nil {
				wrds = []string{""}
				break
			}
			wrds = node.words
		}

		if len(syms) == 0 {
			wrds = root.words
		}
		for i := 0; i < len(wrds); i++ {
			m[wrds[i]] = wrds[i]
		}
	}

	for k := range m {
		ans = append(ans, k)
	}

	sort.Slice(ans, func(i, j int) bool {
		return ans[i] < ans[j]
	})

	return ans
}

func main() {
	scanner := makeScanner()
	data := readArray(scanner)
	req := readArray(scanner)
	ans := solution(data, req)

	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(ans); i++ {
		writer.WriteString(ans[i])
		writer.WriteString("\n")
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

func readArray(scanner *bufio.Scanner) []string {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	arr := make([]string, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		arr[i] = scanner.Text()
	}

	return arr
}
