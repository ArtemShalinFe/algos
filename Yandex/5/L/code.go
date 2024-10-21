package main

func siftDown(heap []int, idx int) int {
	left := 2 * idx
	right := 2*idx + 1

	if left >= len(heap) {
		return idx
	}

	idxLrg := left
	if right < len(heap) && heap[right] > heap[left] {
		idxLrg = right
	}

	if heap[idxLrg] > heap[idx] {
		heap[idx], heap[idxLrg] = heap[idxLrg], heap[idx]
		return siftDown(heap, idxLrg)
	}

	return idx
}

func test() {
	sample := []int{-1, 12, 1, 8, 3, 4, 7}
	if siftDown(sample, 2) != 5 {
		panic("WA")
	}
}
