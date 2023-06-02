package main

import "fmt"

type minHeap struct {
	heap []int
}

func (h *minHeap) parent(i int) int {
	return (i - 1) / 2
}
func (h *minHeap) leftChild(i int) int {
	return (i * 2) + 1
}
func (h *minHeap) rightChild(i int) int {
	return (i * 2) + 2
}

func (h *minHeap) swap(first, second int) {
	h.heap[first], h.heap[second] = h.heap[second], h.heap[first]
}

func (h *minHeap) insert(data int) {
	h.heap = append(h.heap, data)
	i := len(h.heap) - 1
	for i != 0 && h.heap[i] < h.heap[h.parent(i)] {
		h.swap(i, h.parent(i))
	}
}
func (h *minHeap) heapify(i int) {
	smallest := i
	l := h.leftChild(i)
	r := h.rightChild(i)
	if l < len(h.heap) && h.heap[l] < h.heap[smallest] {
		smallest = l
	} 
	if r < len(h.heap) && h.heap[r] < h.heap[smallest] {
		smallest = r
	}
	if smallest != i {
		h.swap(i, smallest)
		h.heapify(smallest)
	}
	
}
func (h *minHeap) extractMin() int {
	root := h.heap[0]
	h.heap[0] = h.heap[len(h.heap)-1]
	h.heap = h.heap[:len(h.heap)-1]
	h.heapify(0)
	return root
}

func main() {
	h := &minHeap{}
	h.insert(3)
	h.insert(2)
	h.insert(1)
	h.insert(15)
	h.insert(5)
	h.insert(4)
	h.insert(45)
	
	fmt.Println(h.heap) // [1 2 4 15 5 3 45]
	fmt.Println(h.extractMin())
	 // 1
	fmt.Println(h.heap)
}