/*
https://contest.yandex.ru/contest/25597/run-report/116887989/

### Принцип работы

	Программа определяет возможность разделения очков на два подмножества таким образом, чтобы сумма очков в разделенных
	множествах была одинаковая. Для этого она находит возможные подмножества от половины общей суммы множества и, если
	такое множество есть, то значит и другую половину подмножества можно составить из оставшихся элементов и их сумма
	будет равна первому подмножеству.

### Доказательство корректности

	Программа определяет возможность разделения очков с помощью динамического программирования.

	Сначала вычисляет общую сумму баллов (sum) и разделяет ее на 2, если общая сумма баллов получилась нечетная, то разделение
	на два множества с одинаковой суммой невозможно.
	Объявляет массив dp (размером половины от общей суммы) с базовым случаем для пустого множества dp[0] = true.

	На каждой итерации цикла по баллам в массиве dp определяются возможные слагаемые для общей сумму баллов разделенной на 2 (sum/2).

	Если после цикла по баллам последний элемент массива dp равен true, значит, что разделение одной части возможно, а
	значит, что сумма не использованных элементов для второго множества так же будет равна сумме первого.

### Временная сложность

	Временная сложность складывается из цикла по K/2, который обходится N раз, где:
	- N - это количество партий;
	- К - общее количество очков, заработанное в партиях.

	Временная сложность: O(N * K)

### Пространственная сложность

	В памяти программы хранится:
	- массив, в котором К/2 элементов, где К - общее количество очков, заработанное в партиях.

	Пространственная сложность: О(К/2).
*/
package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func solution(sum int, arr []int) bool {
	if sum%2 != 0 {
		return false
	}

	k := sum / 2

	dp := make([]bool, k+1)
	dp[0] = true

	for i := 0; i < len(arr); i++ {
		for j := k; j > 0; j-- {
			if arr[i] <= j && !dp[j] {
				dp[j] = dp[j-arr[i]]
			}
		}
	}

	return dp[k]
}

func main() {
	scanner := makeScanner()
	sum, arr := readArray(scanner)
	print(solution(sum, arr))
}

func readArray(scanner *bufio.Scanner) (int, []int) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	arr := make([]int, n)
	sum := 0
	for j := 0; j < len(listString); j++ {
		v, _ := strconv.Atoi(listString[j])
		arr[j] = v
		sum = sum + v
	}
	return sum, arr
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func print(ans bool) {
	writer := bufio.NewWriter(os.Stdout)
	if ans {
		writer.WriteString("True")
	} else {
		writer.WriteString("False")
	}

	writer.Flush()
}
