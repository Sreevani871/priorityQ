/*
This program contains the implemetntation code for all methods declared in interface.Insertion sort is used to enqueue elements into the PQ
*/
package priorityQ

type Priorityqueue struct { //Priorityqueue struct has fields where Size is PQ size,Max is no.of maximum elements,IntArray is holding elements of PQ,ItemCount is Element coutn of PQ
	Size      int
	Max       int
	IntPq     []int
	ItemCount int
}

/*Method implementation for inserting element into PQ*/

func (N *Priorityqueue) Enqueue(element int) int {
	var i int = 0
	if !N.IsFull() { //Checks whether queue is full or not
		if N.ItemCount == 0 {
			N.IntPq = append(N.IntPq, element) //Inserting first element into queue
			N.ItemCount++                      //Increasing element count
		} else { //The element to be inserted is not the first element comparision of priority takes place
			for i = N.ItemCount - 1; i >= 0; i-- {
				if element > N.IntPq[i] {
					if i == N.ItemCount-1 { //checks whether element to be inserted is greater than the element or not
						N.IntPq = append(N.IntPq, N.IntPq[i])
					} else {
						N.IntPq[i+1] = N.IntPq[i]
					}
				} else {
					break
				}
			}

			if i == len(N.IntPq)-1 {
				N.IntPq = append(N.IntPq, element)
				N.ItemCount++
			} else {
				N.IntPq[i+1] = element //Inserting element at right place in PQ
				N.ItemCount++
			}
		}

		return element
	} else {

		return -1
	}
}

/* Method implementation for removing element from the queue */

func (N *Priorityqueue) Dequeue() int {
	if !N.IsEmpty() { //Checking PQ is empty or not
		remEle := N.IntPq[N.ItemCount-1] //cathing Removing element from the PQ
		N.IntPq[N.ItemCount-1] = 0       //Removal of element by making value zero
		N.ItemCount--                    //Decreasing element count
		intQ := make([]int, N.ItemCount, 100)
		copy(intQ, N.IntPq[:N.ItemCount])
		N.IntPq = intQ
		return remEle
	} else {
		return -1 //Returns -1 if PQ is empty
	}

}

/* Method implementation for checking whether PQ is empty or not*/

func (N *Priorityqueue) IsEmpty() bool {
	return N.ItemCount == 0
}

/* Method implementation for checking whether PQ is full or not*/

func (N *Priorityqueue) IsFull() bool {
	N.Max = 100
	return N.ItemCount == N.Max
}

/* Method implementation for merging two PQ's*/

func (Pq1 *Priorityqueue) Merge(q *Priorityqueue) {
	Pq2 := q.IntPq
	n := len(Pq2)
	m := len(Pq1.IntPq)

	i := 0
	j := 0
	k := 0

	Pq3 := make([]int, m+n, 100)

	for k < (m + n) {

		if j == n || (i < m && (Pq1.IntPq[i] >= Pq2[j])) {
			Pq3[k] = Pq1.IntPq[i]
			k++
			i++

		} else {
			Pq3[k] = Pq2[j]

			k++
			j++
		}
	}
	Pq1.IntPq = Pq3
	Pq1.Size = len(Pq3)
	Pq1.ItemCount = len(Pq3)

}
func New() *Priorityqueue {
	return &Priorityqueue{}

}
func (N *Priorityqueue) PrintPQ() []int {
	return N.IntPq
}
