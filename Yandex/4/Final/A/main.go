package main

/*
https://contest.yandex.ru/contest/24414/run-report/115416283/

### Принцип работы

	По всем входным словам документов создается хеш таблица в ключах которой хранится слово, а в значених документы,
	в которых оно встречается. В процессе перебора входных запросов алгоритм находит документы,
	в которых встречаются слова и расчитывает их релевантность. После этого необходимо отсортировать результаты.

### Доказательство корректности

	Для правильного расчета релевантности программа будет строить поисковый индекс на основе входных документов,
	разбивая каждый документ на слова, которые затем будут сохраняться в хеш-таблицу, у которой ключем будет являться
	слово, а значением список документов, в которых это слово встречается.
	При отправке запроса для каждого уникального слова из запроса определяется число его вхождений в документ.
	Полученные числа для всех слов из запроса суммируются - таким образом, поисковая система будет вычислять релевантность
	каждого документа на основе слов в запросе и количества их встречаемости в документах.
	В качестве результата алгоритм возвращает 5 самых релевантных документов.

### Временная сложность

	По условиям задачи количество запросов и документов в базе ограничено 10_000.

	Скорость работы алгоритма зависит от:
	- Количества N документов, используемых для построения поискового индекса и количества слов (K) в этих документах.
	  Максимум из запроса, в котором 1000 символов английского алфавита, можно составить 200 уникальных слов, таким образом
	  в худшем случае индекс может состоять из 10_000*200=2_000_000 слов, но всего в английском языке чуть больше 1_000_000 слов,
	  поэтому будем считать за худший случай 1_000_000 слов в поисковом индексе и 2_000_000 итераций для его построения.
	  Количество слов в тексте документа асимптотически не сильно влияет на скорость построения поискового индекса, их можно опустить - O(N).
	- Количества запросов Q.
	  Максимум из 100 символов английского алфавита можно составить 20 слов, в худшем случае на 1 запрос может быть найдено 20 документов.
	  Количество слов в запросе не сильно влияет на временную сложность и этот показатель можно опустить - O(Q).
	- При формировании ответа на запрос алгоритм в худшем случае найдет L документов.

	Итого:
	- Временная сложность в худшем случае - О(N+(Q*L))

	Про количество слов из символов: https://thecharcounter.com/ru/characters-to-words/
	Про количество слов в англ.языке (ругается на серт): https://languagemonitor.com/category/number-of-words-in-english/new-english-words-per-day/

### Пространственная сложность

	Пространственная сложность зависит от количества данных, которые подаются на вход программе.
	В памяти программы хранится:
	- Массив базы документов для построения поискового индекса размера N.
	- Массив запросов размера Q

	Учитывая оценку возможных объемов из временной сложность можно утверждать, что на пространственную сложность влияет также:
	- Размер построенного поискового индекса M.

	Итого:
	- Пространственная сложность - O(N+Q+M)
*/

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func solution(keys []string, querrys []string) []string {
	var res []string

	idx := buildIndex(keys)

	for i := 0; i < len(querrys); i++ {
		// Для подсчета совпадений документов
		matches := make([]int, len(keys)+1)

		// Для контроля уникальности слов
		uWord := make(map[string]struct{})

		// Расчет релевантности документов
		words := strings.Split(querrys[i], " ")
		for j := 0; j < len(words); j++ {
			word := words[j]

			if _, ok := uWord[word]; ok {
				continue
			}

			uWord[word] = struct{}{}

			if docs, ok := idx[word]; ok {
				for m := 0; m < len(docs); m++ {
					matches[docs[m].id] += 1
				}
			}
		}

		// Формирование результата
		var ans []string
		for k := 0; k < 5; k++ {
			max := 0
			for l := 1; l < len(matches); l++ {
				if matches[max] < matches[l] {
					max = l
				}
			}
			matches[max] = 0
			if max > 0 {
				ans = append(ans, strconv.Itoa(max))
			}
		}
		res = append(res, strings.Join(ans, " "))
	}

	return res
}

type doc struct {
	id   int
	text string
}

func buildIndex(keys []string) map[string][]doc {
	db := make(map[string][]doc, 0)

	for i := 0; i < len(keys); i++ {
		listString := strings.Split(keys[i], " ")
		for j := 0; j < len(listString); j++ {
			word := listString[j]

			if v, ok := db[word]; ok {
				db[word] = append(v, doc{id: i + 1, text: keys[i]})
			} else {
				db[word] = []doc{
					{id: i + 1, text: keys[i]},
				}
			}
		}
	}
	return db
}

func main() {
	scanner := makeScanner()
	ndoc := readInt(scanner)
	keys := readArray(ndoc, scanner)
	qc := readInt(scanner)
	qs := readArray(qc, scanner)
	printRes(solution(keys, qs))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 64 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func printRes(r []string) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(r); i++ {
		writer.WriteString(r[i])
		writer.WriteString("\n")
	}
	writer.Flush()
}

func readArray(n int, scanner *bufio.Scanner) []string {
	var arr []string
	for i := 0; i < n; i++ {
		scanner.Scan()
		arr = append(arr, scanner.Text())
	}
	return arr
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}
