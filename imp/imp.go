/*
This program contains the implemetntation code for all methods declared in interface.
*/
package imp
import(
		"fmt"
)


type Priorityqueue struct{		//Priorityqueue struct has fields where Size is PQ size,Max is no.of maximum elements,IntArray is holding elements of PQ,ItemCount is Element coutn of PQ
  Size int
  Max int
  IntArray []int
  ItemCount int
  }

/*
This New() method generates PQ of given size and asks the user to choose between operations need to perform on PQ
*/

func(p *Priorityqueue) New() int{
		fmt.Println("Enter the queue size: ")
		var a int
		fmt.Scanf("%d",&a)										//Taking size from the 							
		p.Size=a    											//Assigning size to PQ
		p.Max=p.Size
		p.IntArray=make([]int,p.Size)							//Generating PQ of given size
		p.ItemCount=0											//Initially making  no of elements in a queue is zero
		var option int
		fmt.Println("Options\n1)Enqueue\n2)Dequeue\n3)IsEmpty\n4)IsFull\n5)View Queue\n")	 //Display the operations which performs on PQ
		fmt.Println("Choose option :")
		fmt.Scanf("%d",&option)									//Choosing operation to be performed on PQ
		for true{
				if option==1{
							var Data int
							fmt.Println("Enter the element into the queue:")
							fmt.Scanf("%d",&Data)
							p.Enqueue(Data)						//Which insert element into PQ
				}else if option==2{
							p.Dequeue()							//Which remove element from PQ
				}else if option==3{
							Empty:=p.IsEmpty()					//Checking whether PQ is empty or not
							if Empty{
								fmt.Println("Queue is empty")
							}else{
								fmt.Println("Queue is not empty")
							}
				}else if option==4{
							Full:=p.IsFull()					//Checking whether PQ is full or not
							if Full{
								fmt.Println("Queue is Full")
							}else{
								fmt.Println("Queue is not Full")
							}
				}else if option==5{
							p.ViewPriorityQueue()				//Displays PQ
				
				}else{
					break
				}
				fmt.Println("Enter option/(>5) for exit :")					//Choose again an operation need to perform
				fmt.Scanf("%d",&option)
		}
		return p.Size 											//Returns Size of PQ
	}
		
/*Method implementation for inserting element into PQ*/

func (N *Priorityqueue) Enqueue(element int){
	var i int=0
	if !N.IsFull(){												//Checks whether queue is full or not
		if N.ItemCount==0{										//Inserting first element into queue
			N.IntArray[N.ItemCount]=element
			N.ItemCount++										//Increasing element count
		}else{													//The element to be inserted is not the first element comparision of priority takes place
			for i=N.ItemCount-1;i>=0;i--{
				if element >N.IntArray[i]{						//checks whether element to be inserted is greater than the element or not
					N.IntArray[i+1]=N.IntArray[i]				//Shift of element takes place untill inserted elemets finds a right place to insert
				}else{
					break
				}
			}
			N.IntArray[i+1]=element 							//Inserting element at right place in PQ
			N.ItemCount++
		}
	}else{
	fmt.Println("Queue is full")								//Displays a message if PQ is full
	}
}

/* Method implementation for removing element from the queue */
					
func (N *Priorityqueue) Dequeue() int{
	if !N.IsEmpty(){											//Checking PQ is empty or not
	remEle:=N.IntArray[N.ItemCount-1]							//cathing Removing element from the PQ	
	N.IntArray[N.ItemCount-1]=0									//Removal of element by making value zero							
	N.ItemCount--												//Decreasing element count											
	return remEle
	}else{
		fmt.Println("Queue is emty")							//Displays a message if PQ is empty
		return -1
	}
	
}

/* Method implementation for checking whether PQ is empty or not*/

func (N *Priorityqueue) IsEmpty() bool{
	return N.ItemCount==0
}

/* Method implementation for checking whether PQ is full or not*/

func (N *Priorityqueue) IsFull() bool{
	return N.ItemCount==N.Max }

/* Method for dispaying PQ*/

func (N *Priorityqueue) ViewPriorityQueue(){
	fmt.Println(N.IntArray)
	}
