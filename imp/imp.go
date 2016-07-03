/*
This program contains the implemetntation code for all methods declared in interface.
*/
package imp

import (
	"fmt"
	//	P "github.com/Sreevani871/queue1/new"
)

/*Method implementation for inserting element into PQ*/
type Priorityqueue struct { //Priorityqueue struct has fields where Size is PQ size,Max is no.of maximum elements,IntArray is holding elements of PQ,ItemCount is Element coutn of PQ
	Size      int
	Max       int
	IntPq     []int
	ItemCount int
}

func (N *Priorityqueue) Enqueue(element int) {
	var i int = 0
	if !N.IsFull() { //Checks whether queue is full or not
		if N.ItemCount == 0 { //Inserting first element into queue
			N.IntPq[N.ItemCount] = element
			N.ItemCount++ //Increasing element count
		} else { //The element to be inserted is not the first element comparision of priority takes place
			for i = N.ItemCount - 1; i >= 0; i-- {
				if element > N.IntPq[i] { //checks whether element to be inserted is greater than the element or not
					N.IntPq[i+1] = N.IntPq[i] //Shift of element takes place untill inserted elemets finds a right place to insert
				} else {
					break
				}
			}

			N.IntPq[i+1] = element //Inserting element at right place in PQ
			N.ItemCount++
		}
	} else {
		fmt.Println("Queue is full") //Displays a message if PQ is full
	}
}

/* Method implementation for removing element from the queue */

func (N *Priorityqueue) Dequeue() int {
	if !N.IsEmpty() { //Checking PQ is empty or not
		remEle := N.IntPq[N.ItemCount-1] //cathing Removing element from the PQ
		N.IntPq[N.ItemCount-1] = 0       //Removal of element by making value zero
		N.ItemCount--                    //Decreasing element count
		return remEle
	} else {
		fmt.Println("Queue is emty") //Displays a message if PQ is empty
		return -1
	}

}

/* Method implementation for checking whether PQ is empty or not*/

func (N *Priorityqueue) IsEmpty() bool {
	return N.ItemCount == 0
}

/* Method implementation for checking whether PQ is full or not*/

func (N *Priorityqueue) IsFull() bool {
	return N.ItemCount == N.Max
}

/* Method for dispaying PQ*/

func (N *Priorityqueue) ViewPriorityQueue() {
	fmt.Println(N.IntPq)
}
