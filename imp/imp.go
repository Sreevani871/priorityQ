package imp
import(
		"fmt"
)
var N Priorityqueue
func(p *Priorityqueue) New() int{
		fmt.Println("Enter the queue size: ")
		fmt.Scanf("%d",&p.Size)
		p.Max=p.Size
		fmt.Println("Size of the queue is ",p.Size)
		p.IntArray=make([]int,p.Size)
		p.ItemCount=0
		var option int
		fmt.Println("Options\n1)Enqueue\n2)Dequeue\n3)isEmpty\n4)isFull\n5)View Queue\n6)Size\n")
		fmt.Println("Choose option")
		fmt.Scanf("%d",&option)
		for true{
				if option==1{
							var Data int
							fmt.Println("Enter the element into the queue:")
							fmt.Scanf("%d",&Data)
							Enqueue(Data)
				}else if option==2{
							Dequeue()
				}else if option==3{
							Empty:=IsEmpty()
							if Empty{
								fmt.Println("Queue is empty")
							}else{
								fmt.Println("Queue is not empty")
							}
				}else if option==4{
							Full:=IsFull()
							if Full{
								fmt.Println("Queue is Full")
							}else{
								fmt.Println("Queue is not Full")
							}
				}else if option==5{
							ViewPriorityQueue()
				
				}else{
					break
				}
				fmt.Println("Enter ur option")
				fmt.Scanf("%d",&option)
		}
		return p.Size
	}
		

func IsFull() bool{
	return N.ItemCount==N.Max
}


func IsEmpty() bool{
	return N.ItemCount==0
}


func Enqueue(element int){
	var i int=0
	if !IsFull(){
		if N.ItemCount==0{
			N.IntArray[N.ItemCount]=element
			N.ItemCount++
		}else{
			for i=N.ItemCount-1;i>=0;i--{
				if element >N.IntArray[i]{
					N.IntArray[i+1]=N.IntArray[i]
					fmt.Println("iteration:",N.IntArray)
				}else{
					break
				}
			}
			N.IntArray[i+1]=element
			N.ItemCount++
		}
	}else{
	fmt.Println("Queue is full")
	}
}

					
func Dequeue() int{
	if !IsEmpty(){
	remEle:=N.IntArray[N.ItemCount-1]	
	N.IntArray[N.ItemCount-1]=0
	N.ItemCount--
	return remEle
	}else{
		fmt.Println("Queue is emty")
		return -1
	}
	
}
func ViewPriorityQueue(){
	fmt.Println(N.IntArray)
	}
