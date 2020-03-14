package memory

import (
	"errors"
)

// Minheap is an auxialiry data structure to get the closest slot
// in logarithmic time.
type Minheap struct {
	heapArray []int
	size      int
	maxsize   int
}

// ErrHeapFull ...
var ErrHeapFull = errors.New("Heap is full")

// ErrHeapEmpty ...
var ErrHeapEmpty = errors.New("Heap is empty")

// NewMinHeap is a constructor function
func NewMinHeap(n int) *Minheap {
	minheap := &Minheap{
		heapArray: []int{},
		size:      0,
		maxsize:   n,
	}
	for i := 1; i <= n; i++ {
		minheap.Insert(i)
	}
	minheap.buildMinHeap()
	return minheap
}

// Insert a new int into the heap.
func (m *Minheap) Insert(item int) error {
	if m.size >= m.maxsize {
		return ErrHeapFull
	}
	m.heapArray = append(m.heapArray, item)
	m.size++
	m.upHeapify(m.size - 1)
	return nil
}

// GetNearestSlot gives the smallest value in the heap.
func (m *Minheap) GetNearestSlot() (int, error) {
	if m.size == 0 {
		return -1, ErrHeapEmpty
	}
	top := m.heapArray[0]
	m.heapArray[0] = m.heapArray[m.size-1]
	m.heapArray = m.heapArray[:(m.size)-1]
	m.size--
	m.downHeapify(0)
	return top, nil
}

func (m *Minheap) leaf(index int) bool {
	if index >= (m.size/2) && index <= m.size {
		return true
	}
	return false
}

func (m *Minheap) parent(index int) int {
	return (index - 1) / 2
}

func (m *Minheap) leftchild(index int) int {
	return 2*index + 1
}

func (m *Minheap) rightchild(index int) int {
	return 2*index + 2
}

func (m *Minheap) swap(first, second int) {
	temp := m.heapArray[first]
	m.heapArray[first] = m.heapArray[second]
	m.heapArray[second] = temp
}

func (m *Minheap) upHeapify(index int) {
	for m.heapArray[index] < m.heapArray[m.parent(index)] {
		m.swap(index, m.parent(index))
	}
}

func (m *Minheap) downHeapify(current int) {
	if m.leaf(current) {
		return
	}
	smallest := current
	leftChildIndex := m.leftchild(current)
	rightRightIndex := m.rightchild(current)
	//If current is smallest then return
	if leftChildIndex < m.size && m.heapArray[leftChildIndex] < m.heapArray[smallest] {
		smallest = leftChildIndex
	}
	if rightRightIndex < m.size && m.heapArray[rightRightIndex] < m.heapArray[smallest] {
		smallest = rightRightIndex
	}
	if smallest != current {
		m.swap(current, smallest)
		m.downHeapify(smallest)
	}
	return
}

func (m *Minheap) buildMinHeap() {
	for index := ((m.size / 2) - 1); index >= 0; index-- {
		m.downHeapify(index)
	}
}
