package main

import "fmt"

// MaxHeap is a struct that holds a slice to store data
type MaxHeap struct {
	array []int
}

// Insert method adds an element to heap
func (h *MaxHeap) Insert(key int) {
	// adding the key to slice
	h.array = append(h.array, key)
	// heapify starts from the last elemnet inserted into array
	h.HeapUp(len(h.array) - 1)
}

// Extract method returns the largest key and removes it from the heap
func (h *MaxHeap) Extrat() int {
	largest := h.array[0]
	lastIndex := len(h.array) - 1
	if len(h.array) == 1 {
		return -1
	}
	h.array[0] = h.array[lastIndex]
	h.array = h.array[:lastIndex] //removing the duplicate last key
	h.HeapDown(0)                 //starting from the largest element

	return largest
}

func (h *MaxHeap) HeapUp(index int) {
	for h.array[getParentIndex(index)] < h.array[index] {
		h.swap(getParentIndex(index), index)
		index = getParentIndex(index)
	}
}

func (h *MaxHeap) HeapDown(index int) {
	lastIndex := len(h.array) - 1
	left, right := getLeftChildIndex(index), getRightChildIndex(index)
	childToCompare := 0
	//current index has atleast one child node
	for left <= lastIndex {
		// when the current index has only left child
		if left == lastIndex {
			childToCompare = left
		} else if h.array[left] > h.array[right] { // when left child is greater than the right child
			childToCompare = left
		} else { // when right child is greater than the left child
			childToCompare = right
		}
		// if the value at current index is less than the childToCompare we need to swap
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			//update the index and left,right childs
			index := childToCompare
			left, right = getLeftChildIndex(index), getRightChildIndex(index)
		} else {
			return
		}
	}
}

//get parent node index
func getParentIndex(i int) int {
	return (i - 1) / 2
}

//get the left child index
func getLeftChildIndex(i int) int {
	return 2*i + 1
}

//get the right child index
func getRightChildIndex(i int) int {
	return 2*i + 2
}

//swap nodes
func (h *MaxHeap) swap(a, b int) {
	h.array[a], h.array[b] = h.array[b], h.array[a]
}

func main() {
	maxheap := &MaxHeap{}
	constructHeap := []int{1, 2, 4, 5, 7, 8, 3}
	for _, v := range constructHeap {
		maxheap.Insert(v)
		fmt.Println("Insert", maxheap)
	}
	for i := 0; i < 5; i++ {
		largest := maxheap.Extrat()
		fmt.Println("Extract", maxheap, largest)
	}
}
