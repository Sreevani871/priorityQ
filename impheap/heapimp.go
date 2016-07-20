/*Priority Queue implementation using Min-Heap */

package priorityQ

import "sort"

//import "math/rand"

/* MinHeap struct is created with appropriate fields*/

type Heap struct {
	heapSlice []int
	length    int
	max       int
}

/* Method Min_heapify ensure heap is statisfying min heap property or not */

func (H *Heap) Min_heapify(i int, N int) {

	left := 2 * i
	right := 2*i + 1
	var smallest int
	if left <= N && H.heapSlice[left] < H.heapSlice[i] {
		smallest = left
	} else {
		smallest = i
	}
	if right <= N && H.heapSlice[right] < H.heapSlice[smallest] {
		smallest = right
	}
	if smallest != i {
		H.heapSlice[i], H.heapSlice[smallest] = H.heapSlice[smallest], H.heapSlice[i]
		H.Min_heapify(smallest, N)
	}
}

/* Minimum method returns minimum interger from the heap */

func (H *Heap) Minimum() int {
	return H.heapSlice[1]
}

/* Method for removing elements from heap */

func (H *Heap) Dequeue() int {
	if H.length == 0 {
		H.heapSlice = make([]int, 2, 100)
		return -1
	}
	min := H.heapSlice[1]
	H.heapSlice[1] = H.heapSlice[H.length]
	H.length = H.length - 1
	H.Min_heapify(1, H.length)

	H.heapSlice = H.heapSlice[:H.length+1]

	return min
}

/* Decrease_value method is used by Enqueue method for inserting elements into the heap */

func (H *Heap) Decrease_value(i int, val int) int {

	if i == 1 {
		H.heapSlice[i] = val
	} else {
		H.heapSlice = append(H.heapSlice, val)

	}

	for i > 1 && H.heapSlice[i/2] > H.heapSlice[i] {
		H.heapSlice[i/2], H.heapSlice[i] = H.heapSlice[i], H.heapSlice[i/2]
		i = i / 2
	}
	return val

}

/* Method for inserting elements into the queuw */

func (H *Heap) Enqueue(val int) int {
	if H.length == 0 {
		H.heapSlice = make([]int, 2, 100)
		H.max = 100
	}
	if !H.IsFull() {
		H.length = H.length + 1

		return H.Decrease_value(H.length, val)
	} else {
		return -1
	}
}

/* Method for checking whether heap is empty or not */

func (H *Heap) IsEmpty() bool {
	return H.length == 0
}

/* Method for checking whether heap is full or not */

func (H *Heap) IsFull() bool {
	H.max = 100
	return H.length == H.max
}

/* Method for merging two heaps */

func (Pq4 *Heap) Merge(q *Heap) {
	Pq2 := q.heapSlice[1:]
	n := len(Pq2)
	m := len(Pq4.heapSlice)
	//fmt.Println("Not there here is actual problem")
	Pq1 := make([]int, m-1, len(Pq4.heapSlice)+10)
	//fmt.Println("fine")
	copy(Pq1, Pq4.heapSlice[1:])
	sort.Ints(Pq2)
	sort.Ints(Pq1)
	//fmt.Println("fine1")
	m = m - 1

	i := 0
	j := 0
	k := 0
	Pq3 := make([]int, m+n, m+n)
	//fmt.Println("fine2")
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
	//fmt.Println("Here is the problem")
	Pq4.heapSlice = make([]int, len(Pq3)+1, len(Pq3)+10)
	copy(Pq4.heapSlice[1:], Pq3[:])

	Pq4.length = len(Pq3) - 1
}

/* New() method returns struct of heap to implement all methods in interface */

func New() *Heap {

	return &Heap{}
}

/* Method for dispaying heap */

func (H *Heap) PrintPQ() []int {
	return H.heapSlice
}
