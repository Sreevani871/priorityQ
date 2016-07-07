package priorityQ

import "fmt"

import "sort"

type Heap struct {
	HeapSlice []int
	Length    int
	Max       int
}

var C Heap

func (H *Heap) Min_heapify(i int, N int) {

	left := 2 * i
	right := 2*i + 1
	var smallest int
	if left <= N && H.HeapSlice[left] < H.HeapSlice[i] {
		smallest = left
	} else {
		smallest = i
	}
	if right <= N && H.HeapSlice[right] < H.HeapSlice[smallest] {
		smallest = right
	}
	if smallest != i {
		H.HeapSlice[i], H.HeapSlice[smallest] = H.HeapSlice[smallest], H.HeapSlice[i]
		H.Min_heapify(smallest, N)
	}
}

func (H *Heap) Minimum() int {
	return H.HeapSlice[1]
}

func (H *Heap) Dequeue() int {
	if H.Length == 0 {
		H.HeapSlice = make([]int, 2, 100)
		return -1
	}
	min := H.HeapSlice[1]
	H.HeapSlice[1] = H.HeapSlice[H.Length]
	H.Length = H.Length - 1
	H.Min_heapify(1, H.Length)

	H.HeapSlice = H.HeapSlice[:H.Length+1]

	return min
}
func (H *Heap) Decrease_value(i int, val int) int {

	if i == 1 {
		H.HeapSlice[i] = val
	} else {
		H.HeapSlice = append(H.HeapSlice, val)

	}

	for i > 1 && H.HeapSlice[i/2] > H.HeapSlice[i] {
		H.HeapSlice[i/2], H.HeapSlice[i] = H.HeapSlice[i], H.HeapSlice[i/2]
		i = i / 2
	}
	return val

}
func (H *Heap) Enqueue(val int) int {
	if H.Length == 0 {
		H.HeapSlice = make([]int, 2, 100)
		H.Max = 100
	}
	if !H.IsFull() {
		H.Length = H.Length + 1

		return H.Decrease_value(H.Length, val)
	} else {
		return -1
	}
}
func (H *Heap) IsEmpty() bool {
	return H.Length == 0
}

func (H *Heap) IsFull() bool {
	H.Max = 100
	return H.Length == H.Max
}
func (Pq4 *Heap) Merge(q *Heap) {
	Pq2 := q.HeapSlice[1:]
	n := len(Pq2)
	m := len(Pq4.HeapSlice)
	Pq1 := make([]int, m-1, 100)
	copy(Pq1, Pq4.HeapSlice[1:])
	fmt.Println("two heaps:", Pq1, Pq2)
	sort.Ints(Pq2)
	sort.Ints(Pq1)
	fmt.Println(Pq2, Pq1)
	m = m - 1

	i := 0
	j := 0
	k := 0
	Pq3 := make([]int, m+n, 100)

	for k < (m + n) {

		if j == n || (i < m && (Pq1[i] <= Pq2[j])) {
			Pq3[k] = Pq1[i]
			k++
			i++

		} else {
			Pq3[k] = Pq2[j]

			k++
			j++
		}
	}
	Pq4.HeapSlice = make([]int, len(Pq3)+1, 100)
	copy(Pq4.HeapSlice[1:], Pq3[:])

	Pq4.Length = len(Pq3) - 1
}

func New() *Heap {

	return &Heap{}
}

func (H *Heap) PrintPQ() []int {
	return H.HeapSlice
}
