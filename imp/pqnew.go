/*
This program contains an interface having methods which impelemts all the basic queue operations 
*/
package imp
type queue interface{
	New() int 
	Enqueue(e int)
	Dequeue() int
	IsFull() bool
	IsEmpty()bool
	ViewPriorityQueue()
}







