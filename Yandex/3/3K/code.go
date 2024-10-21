package main

import "reflect"

func merge_sort(arr []int, lf int, rg int) {
	if len(arr) == 1 {
		return
	}
	m := (lf + rg) / 2
	arrLeft := arr[:m]
	merge_sort(arrLeft, lf, m)

	arrRight := arr[m:]
	merge_sort(arrRight, 0, len(arrRight))

	res := merge(arr, lf, m, len(arrRight))
	copy(arr, res)
}

func merge(arr []int, left int, mid int, right int) (result []int) {
	var res []int

	arrLeft := arr[:mid]
	arrRight := arr[mid:]

	arrL := 0
	arrR := 0
	for {
		if arrLeft[arrL] < arrRight[arrR] {
			res = append(res, arrLeft[arrL])
			arrL++
		} else {
			res = append(res, arrRight[arrR])
			arrR++
		}

		if arrL >= len(arrLeft) {
			for i := arrR; i < len(arrRight); i++ {
				res = append(res, arrRight[i])
			}
			break
		}

		if arrR >= len(arrRight) {
			for i := arrL; i < len(arrLeft); i++ {
				res = append(res, arrLeft[i])
			}
			break
		}
	}
	return res
}

func main() {
	a := []int{1, 4, 9, 2, 10, 11}
	b := merge(a, 0, 3, 6)
	expected := []int{1, 2, 4, 9, 10, 11}
	if !reflect.DeepEqual(b, expected) {
		panic("WA. Merge")
	}

	c := []int{1, 4, 2, 10, 1, 2}
	merge_sort(c, 0, 6)
	expected = []int{1, 1, 2, 2, 4, 10}
	if !reflect.DeepEqual(c, expected) {
		panic("WA. MergeSort")
	}
}
