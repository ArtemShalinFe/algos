package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func solution(a int, m int, s string) int {
	l := len(s)

	if l == 0 {
		return 0
	}

	if l == 1 {
		return int(s[0]) % m
	}

	result := ((int(s[0]) * a) + int(s[1])) % m
	for i := 2; i < l; i++ {
		result = ((result * a) + int(s[i])) % m
	}
	return result
}

func main() {
	mp := make(map[int]string)
	a := 1000
	m := 123987123

	anss := make([]string, 0)
	for len(anss) < 2 {
		s := genString()
		hash := solution(a, m, s)
		v, ok := mp[hash]
		if ok {
			b := solution(a, m, v)
			if v != s && b == hash {
				anss = append(anss, s)
			}
		} else {
			mp[hash] = s
		}
	}

	for i := 0; i < len(anss); i++ {
		fmt.Printf("%s --- %d\n", anss[i], solution(a, m, anss[i]))
	}
}

func genString() string {
	b := make([]rune, 500)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

var letters = []rune("abcdefghijklmnopqrstuvwxyz")

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func printRes(m int) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(strconv.Itoa(m))
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
