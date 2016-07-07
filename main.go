package main

import (
	"fmt"
	S2 "github.com/Sreevani871/priorityQ/imp"
	S1 "github.com/Sreevani871/priorityQ/impheap"
)

type Slice struct {
	Array []int
}

func main() {
	var Option int
	fmt.Println("Choose Optinon\n1)PriorityQ by insertion sort\n2)PriorityQ by minHeap\n")
	fmt.Scanf("%d", &Option)
	N := S2.New()                                                                         // New() method returns a struct of PriorityQ
	H := S1.New()                                                                         // New() method returns a struct of PriorityQ
	fmt.Println("Options\n1)Enqueue\n2)Dequeue\n3)IsEmpty\n4)IsFull\n5)Merge\n6)PrintPQ") //Display the operations which performs on PQ
	fmt.Println("Choose operation :")
	var op int
	fmt.Scanf("%d", &op)
	for true {
		if op == 1 {
			var Data int
			fmt.Println("Enter the element into the queue:")
			fmt.Scanf("%d", &Data)
			if Option == 1 {
				N.Enqueue(Data) //Inserting elements into the PQ
			}
			if Option == 2 {
				H.Enqueue(Data) //Inserting elements into the PQ
			}
		} else if op == 2 {
			if Option == 1 {
				fmt.Println(N.Dequeue()) //Removing element from the PQ
			}
			if Option == 2 {
				fmt.Println(H.Dequeue()) //Removing element from the PQ
			}
		} else if op == 3 {
			var Empty bool
			if Option == 1 {
				Empty = N.IsEmpty() //Checking whether PQ is empty or not
			}
			if Option == 2 {
				Empty = H.IsEmpty() //Checking whether PQ is empty or not
			}
			if Empty {
				fmt.Println("Queue is empty")
			} else {
				fmt.Println("Queue is not empty")
			}
		} else if op == 4 {
			var Full bool
			if Option == 2 {
				Full = N.IsFull() //Checking whether PQ is full or not
			}
			if Option == 1 {
				Full = H.IsFull() //Checking whether PQ is full or not
			}
			if Full {
				fmt.Println("Queue is Full")
			} else {
				fmt.Println("Queue is not Full")
			}
		} else if op == 5 {
			P := S2.New()
			P.Enqueue(91)
			P.Enqueue(0)
			P.Enqueue(-90)
			P.Enqueue(1000)
			if Option == 1 {
				N.Merge(P)
			} else if Option == 2 {
				Q := S1.New()
				Q.Enqueue(6)
				Q.Enqueue(0)
				Q.Enqueue(3)
				Q.Enqueue(1)
				Q.Enqueue(-90)
				Q.Enqueue(67)
				H.Merge(Q)

			}

		} else if op == 6 {
			if Option == 1 {
				fmt.Println("PQ:", N.PrintPQ())
			} else if Option == 2 {
				fmt.Println("PQ:", H.PrintPQ())
			}

		} else {
			break
		}
		fmt.Println("Enter operation/(>6) for exit :") //Choose again an operation need to perform
		fmt.Scanf("%d", &op)
	}
}
