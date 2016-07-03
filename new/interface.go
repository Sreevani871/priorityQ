package new

type PqInterface interface {
	Enqueue(int) int
	Dequeue() int
	IsEmpty() bool
	IsFull() bool
	ViewPriorityQueue()
}
