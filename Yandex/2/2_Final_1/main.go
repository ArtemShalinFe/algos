/*
-- ССЫЛКА НА ОТЧЕТ --
https://contest.yandex.ru/contest/22781/run-report/114551369/

-- ПРИНЦИП РАБОТЫ --
Я реализовал кольцевой буфер используя слайс с заранее определенным (maxSize) количеством элементов.
Используя переменные headIdx, tailIdx для хранения текущих указателей на индексы "головы и хвоста" списка элементов, могу
за постоянное время добавлять и извлекать элементы.
Переменная size используется, чтобы контролировать текущий размер слайса.

Элементы могут добавляться как в начало, так и в конец списка элементов, при этом:
- Перед добавлением элемента в начало слайса к индексу headIdx добавляется 1.
- После добавлением элемента в конец слайса из индекса tailIdx вычитается 1.

В обоих случаях добавление происходит за постоянное время так как заранее известно в какую ячейку слайса нужно поместить
элемент, а переменная size инкрементируется на 1.

Элементы могут извлекаться как с начала дека, так и с конца.
В случае извлечения элементов переменные headIdx, tailIdx меняют свои значения на обратные при добавлении:
- Перед извлечением элемента с начала слайса из индексу headIdx вычитается 1.
- После добавления элемента с конца слайса к индексу tailIdx добавляется 1.

В обоих случаях значение извлекается за постоянное время так как заранее известно какой индекс у первого и последнего
элемента, а переменная size декрементируется на 1.

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
Из описания алгоритма следует, что кольцевой буфер реализован на слайсе с заранее указанным количеством элементов,
поэтому при добавлении в него новых элементов не будет происходит реаллокация слайса. С помощью переменных headIdx, tailIdx
алгоритму заранее известны индексы первого и последнего элемента для обработки.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
Время работы алгоритма зависит от количества команд, которые пришли во входных данных, поэтому временная сложность O(n).

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
Под дек алгоритму необходимо выделить память равную размеру переданной в алгоритм переменной n,
поэтому дек будет потреблять O(n) памяти.

Дополнительно алгоритму требуется выделить O(k) памяти, под входные данные.

Итого пространственная сложность алгоритма О(n+k) памяти.
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := makeScanner()
	cmdCount := readInt(scanner)
	dequeSize := readInt(scanner)
	cmds := readCmds(cmdCount, scanner)
	printArray(solution(dequeSize, cmds))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func printArray(ans []string) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(ans); i++ {
		writer.WriteString(ans[i])
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

func readCmds(cmdCount int, scanner *bufio.Scanner) []string {
	cmds := make([]string, cmdCount)

	for i := 0; i < cmdCount; i++ {
		scanner.Scan()
		cmds[i] = scanner.Text()
	}

	return cmds
}

func solution(dequeSize int, cmds []string) []string {
	ans := make([]string, 0)

	d := NewDeque(dequeSize)

	for i := 0; i < len(cmds); i++ {
		cmd := cmds[i]

		c := strings.Split(cmd, " ")
		if len(c) < 1 {
			log.Fatal("unexpected cmd format")
		}

		action := c[0]

		switch {
		case action == pushBack:
			if d.isFull() {
				ans = append(ans, errs)
				continue
			}

			number := c[1]
			num, err := d.parseNum(number)
			if err != nil {
				log.Fatalf("an error occured while parse number in cmd(%s), err: %v", cmd, err)
			}

			d.pushBack(num)
		case action == pushFront:
			if d.isFull() {
				ans = append(ans, errs)
				continue
			}

			number := c[1]
			num, err := d.parseNum(number)
			if err != nil {
				log.Fatalf("an error occured while parse number in cmd(%s), err: %v", cmd, err)
			}

			d.pushFront(num)
		case action == popBack:
			if d.isEmpty() {
				ans = append(ans, errs)
				continue
			}

			ans = append(ans, strconv.Itoa(d.popBack()))
		case action == popFront:
			if d.isEmpty() {
				ans = append(ans, errs)
				continue
			}

			ans = append(ans, strconv.Itoa(d.popFront()))
		default:
			continue
		}
	}

	return ans
}

const (
	errs      = "error"
	pushBack  = "push_back"
	popBack   = "pop_back"
	pushFront = "push_front"
	popFront  = "pop_front"
)

type Deque struct {
	queue   []int
	headIdx int
	tailIdx int
	size    int
	maxSize int
}

func NewDeque(dequeSize int) *Deque {
	return &Deque{
		queue:   make([]int, dequeSize),
		headIdx: 0,
		tailIdx: 0,
		size:    0,
		maxSize: dequeSize,
	}
}

func (d *Deque) pushFront(x int) {
	d.queue[d.headIdx] = x
	d.moveHeadIdx(1)

	d.size++
}

func (d *Deque) pushBack(x int) {
	d.moveTailIdx(-1)
	d.queue[d.tailIdx] = x

	d.size++
}

func (d *Deque) popFront() int {
	d.moveHeadIdx(-1)
	value := d.queue[d.headIdx]
	d.queue[d.headIdx] = 0
	d.size--

	return value
}

func (d *Deque) popBack() int {
	value := d.queue[d.tailIdx]
	d.queue[d.tailIdx] = 0
	d.moveTailIdx(1)
	d.size--

	return value
}

func (d *Deque) moveHeadIdx(dirrection int) {
	d.headIdx = (d.headIdx + dirrection) % d.maxSize
	if d.headIdx < 0 {
		d.headIdx = d.maxSize - 1
	}
}

func (d *Deque) moveTailIdx(dirrection int) {
	d.tailIdx = (d.tailIdx + dirrection) % d.maxSize
	if d.tailIdx < 0 {
		d.tailIdx = d.maxSize - 1
	}
}

func (d *Deque) isFull() bool {
	return d.size == d.maxSize
}

func (d *Deque) isEmpty() bool {
	return d.size == 0
}

func (d *Deque) parseNum(nmb string) (int, error) {
	num, err := strconv.Atoi(nmb)
	if err != nil {
		return 0, fmt.Errorf("an error occured while parse value to int, err: %w", err)
	}

	return num, nil
}
