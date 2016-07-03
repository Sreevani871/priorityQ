package new

import (
	"fmt"
	S "github.com/Sreevani871/queue1/imp"
)

var P S.Priorityqueue

/*
This New() Function generates PQ of given size and asks the user to choose between operations need to perform on PQ.And we can increase or decrease the size of PQ
*/

func New() int {
	fmt.Println("Enter the queue size: ")
	var a int
	fmt.Scanf("%d", &a) //Taking size from the
	P.Size = a          //Assigning size to PQ
	P.Max = P.Size
	P.IntPq = make([]int, a, 100) //Generating PQ of given size
	P.ItemCount = 0               //Initially making  no of elements in a queue is zero
	var option int
	fmt.Println("Options\n1)Enqueue\n2)Dequeue\n3)IsEmpty\n4)IsFull\n5)View Queue\n6)Grow\n7)Shrink\n") //Display the operations which performs on PQ
	fmt.Println("Choose option :")
	fmt.Scanf("%d", &option) //Choosing operation to be performed on PQ
	for true {
		if option == 1 {
			var Data int
			fmt.Println("Enter the element into the queue:")
			fmt.Scanf("%d", &Data)
			P.Enqueue(Data) //Which insert element into PQ
		} else if option == 2 {
			P.Dequeue() //Which remove element from PQ
		} else if option == 3 {
			Empty := P.IsEmpty() //Checking whether PQ is empty or not
			if Empty {
				fmt.Println("Queue is empty")
			} else {
				fmt.Println("Queue is not empty")
			}
		} else if option == 4 {
			Full := P.IsFull() //Checking whether PQ is full or not
			if Full {
				fmt.Println("Queue is Full")
			} else {
				fmt.Println("Queue is not Full")
			}
		} else if option == 5 {
			P.ViewPriorityQueue() //Displays PQ

		} else if option == 6 { //Can Grow PQ by offset
			var offset int
			fmt.Scanf("%d", &offset)
			slice := make([]int, P.Max+offset, 100)
			copy(slice, P.IntPq)
			P.IntPq = slice
			P.Max = len(P.IntPq)
			//fmt.Println("max:", len(P.IntPq))

		} else if option == 7 { //Can Shrink PQ by limit
			var limit int
			fmt.Scanf("%d", &limit)
			slice := make([]int, limit, 100)
			copy(slice, P.IntPq[:limit])
			P.IntPq = slice
			P.Max = len(P.IntPq)
			//fmt.Println("max:", len(P.IntPq))
			P.ItemCount = len(P.IntPq)

		} else {
			break
		}
		fmt.Println("Enter option(or)(>7) for exit :") //Choose again an operation need to perform
		fmt.Scanf("%d", &option)
	}
	return P.Size //Returns Size of PQ
}
