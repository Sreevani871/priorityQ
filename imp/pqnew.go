package imp

type queue interface{
	New() int 
	Enqueue(e int)
	Dequeue() int
	IsFull() bool
	IsEmpty()bool
	ViewPriorityQueue()
}

type Priorityqueue struct{
  Size int
  Max int
  IntArray []int
  ItemCount int
  }




