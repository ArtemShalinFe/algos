package main

/*
https://contest.yandex.ru/contest/24414/run-report/115417951/

### Принцип работы

	Для хранения зарплаты инициализируется хеш таблица с массивом определенного размера. Значения в массиве являются
	элементами связного списка. То есть для решения коллизий будет использоваться метод открытой адресации.
    Для расчета хеша от id сотрудника используется хеш-функция Дженкинса.
	Полученный хеш в результате будет являться индексом в массиве куда попадет элемент.

### Доказательство корректности

	Хэш-таблица использует массив для хранения записей, которые представлены в виде односвязного списка.
	Односвязный список содержит id сотрудника (ключ хеш таблицы), его зарплату (значение хеш таблицы) и ссылку на следующий элемент (для случая, если произошла коллизия).
	У хеш таблицы есть функции:
	- put будет добавлять новый элемент односвязного списка в массив или, в случае коллизии, добавлять подчиненный элемент в конец односвязного списка.
	- get будет получать значение, которое связано с запрашиваемым ключом. Если элемента нет, то будет выводить "None".
	- delete будет получать значение, связанное с запрашиваемым ключом и удалять элемент односвязного списка по ключу. Если элемента нет, то будет выводить "None".
	Эти функции обеспечивают корректную работу хэш-таблицы, правильно обрабатывая коллизии, удаления и извлечение значений.

### Временная сложность

	Скорость выполнения программы так же зависит от количества N команд, подаваемых на вход программе.
	Операция с хеш таблицой в худшем случае происходит за линейное время K.
	Итого:
	- Временная сложность в худшем случае - O(N*K)
	- Временная сложность в среднем случае - O(N)

### Пространственная сложность

	В памяти программы хранится хеш таблица размера N и массив входных команд K.
	Итого:
	- Пространственная сложность - O(N+K)
*/

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := makeScanner()
	cmdCount := readInt(scanner)
	cmds := readCmds(cmdCount, scanner)
	solution(cmdCount, cmds)
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 128 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}

func readCmds(cmdCount int, scanner *bufio.Scanner) []string {
	cmds := make([]string, cmdCount)

	for i := 0; i < cmdCount; i++ {
		scanner.Scan()
		cmds[i] = scanner.Text()
	}

	return cmds
}

const (
	get    = "get"
	put    = "put"
	delete = "delete"
)

type nodeList struct {
	id     string
	salary int
	next   *nodeList
}

type hashTable struct {
	buckets []*nodeList
	size    int
}

func newHashTable(size int) *hashTable {
	return &hashTable{
		buckets: make([]*nodeList, size),
		size:    size,
	}
}

func (ht *hashTable) get(key string) (int, bool) {
	idx := ht.getIdxWithJenkinsHashFn(key)
	if ht.buckets[idx] != nil {
		startingNode := ht.buckets[idx]
		for {
			if startingNode.id == key {
				return startingNode.salary, true
			}

			if startingNode.next == nil {
				break
			} else {
				startingNode = startingNode.next
			}
		}
	}
	return 0, false
}

func (ht *hashTable) put(key string, value int) {
	idx := ht.getIdxWithJenkinsHashFn(key)

	nd := nodeList{
		id:     key,
		salary: value,
	}

	if ht.buckets[idx] == nil {
		ht.buckets[idx] = &nd
	} else {
		startingNode := ht.buckets[idx]

		for {
			if startingNode.id == key {
				startingNode.salary = value
				return
			}

			if startingNode.next == nil {
				startingNode.next = &nd
				return
			}

			startingNode = startingNode.next
		}
	}
}

func (ht *hashTable) delete(key string) (int, bool) {
	idx := ht.getIdxWithJenkinsHashFn(key)
	if ht.buckets[idx] != nil {
		startingNode := ht.buckets[idx]
		if startingNode.id == key {
			ht.buckets[idx] = startingNode.next
			return startingNode.salary, true
		}

		for {
			if startingNode.next.id == key {
				ans := startingNode.next.salary
				startingNode.next = startingNode.next.next
				return ans, true
			}

			if startingNode.next == nil {
				break
			} else {
				startingNode = startingNode.next
			}
		}
	}

	return 0, false
}

func (ht *hashTable) getIdxWithJenkinsHashFn(key string) int {
	var hash uint32
	for _, ch := range key {
		hash += uint32(ch)
		hash += hash << 10
		hash ^= hash >> 6
	}

	hash += hash << 3
	hash ^= hash >> 11
	hash += hash << 15

	return int(hash) % ht.size
}

func solution(size int, cmds []string) {
	writer := bufio.NewWriter(os.Stdout)

	ht := newHashTable(size)

	for i := 0; i < len(cmds); i++ {
		cmd := cmds[i]

		c := strings.Split(cmd, " ")
		if len(c) < 1 {
			log.Fatal("unexpected cmd format")
		}

		action := c[0]
		key := c[1]

		switch {
		case action == get:
			num, ok := ht.get(key)
			if !ok {
				writer.WriteString("None")
				writer.WriteString("\n")
				writer.Flush()
			} else {
				writer.WriteString(strconv.Itoa(num))
				writer.WriteString("\n")
				writer.Flush()
			}
		case action == put:
			value, _ := strconv.Atoi(c[2])
			ht.put(key, value)
		case action == delete:
			v, ok := ht.delete(key)
			if !ok {
				writer.WriteString("None")
				writer.WriteString("\n")
				writer.Flush()
			} else {
				writer.WriteString(strconv.Itoa(v))
				writer.WriteString("\n")
				writer.Flush()
			}
		default:
			log.Fatal("unexpected cmd format")
		}
	}

}
