package priorityQ

type PqInterface interface {
	Enqueue(int) int
	Dequeue() int
	IsEmpty() bool
	IsFull() bool
	Merge(priorityQ.Priorityqueue)
}
