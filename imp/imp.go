/*
This program contains the implemetntation code for all methods declared in interface.Insertion sort is used to enqueue elements into the PQ
*/
package priorityQ

type Priorityqueue struct { //Priorityqueue struct has fields where Max is no.of maximum elements,IntArray is holding elements of PQ,ItemCount is Element coutn of PQ
	max       int
	intPq     []int
	itemCount int
}

/*Method implementation for inserting element into PQ*/

func (N *Priorityqueue) Enqueue(element int) int {
	var i int = 0
	if !N.IsFull() { //Checks whether queue is full or not
		if N.itemCount == 0 {
			N.intPq = append(N.intPq, element) //Inserting first element into queue
			N.itemCount++                      //Increasing element count
		} else { //The element to be inserted is not the first element comparision of priority takes place
			for i = N.itemCount - 1; i >= 0; i-- {
				if element > N.intPq[i] {
					if i == N.itemCount-1 { //checks whether element to be inserted is greater than the element or not
						N.intPq = append(N.intPq, N.intPq[i])
					} else {
						N.intPq[i+1] = N.intPq[i]
					}
				} else {
					break
				}
			}

			if i == len(N.intPq)-1 {
				N.intPq = append(N.intPq, element)
				N.itemCount++
			} else {
				N.intPq[i+1] = element //Inserting element at right place in PQ
				N.itemCount++
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
		remEle := N.intPq[N.itemCount-1] //cathing Removing element from the PQ
		N.intPq[N.itemCount-1] = 0       //Removal of element by making value zero
		N.itemCount--                    //Decreasing element count
		intQ := make([]int, N.itemCount, 100)
		copy(intQ, N.intPq[:N.itemCount])
		N.intPq = intQ
		return remEle
	} else {
		return -1 //Returns -1 if PQ is empty
	}

}

/* Method implementation for checking whether PQ is empty or not*/

func (N *Priorityqueue) IsEmpty() bool {
	return N.itemCount == 0
}

/* Method implementation for checking whether PQ is full or not*/

func (N *Priorityqueue) IsFull() bool {
	N.max = 100
	return N.itemCount == N.max
}

/* Method implementation for merging two PQ's*/

func (Pq1 *Priorityqueue) Merge(q *Priorityqueue) {
	Pq2 := q.intPq
	n := len(Pq2)
	m := len(Pq1.intPq)

	i := 0
	j := 0
	k := 0

	Pq3 := make([]int, m+n, 100)

	for k < (m + n) {

		if j == n || (i < m && (Pq1.intPq[i] >= Pq2[j])) {
			Pq3[k] = Pq1.intPq[i]
			k++
			i++

		} else {
			Pq3[k] = Pq2[j]

			k++
			j++
		}
	}
	Pq1.intPq = Pq3
	Pq1.itemCount = len(Pq3)

}
func New() *Priorityqueue {
	return &Priorityqueue{}

}
func (N *Priorityqueue) PrintPQ() []int {
	return N.intPq
}
