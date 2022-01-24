package internal

import (
	"errors"
	"sync"
)

type Queue struct {
	Front int
	Rear int
	Capacity int
	Elements []int
	Mutex sync.Mutex
}

func New(capacity int) *Queue{
	//initialize array with size
	queue:= Queue{
		Elements: make([]int,capacity),
		Capacity: capacity,
		Rear: -1,
		Front: -1,
	}
	return &queue
}


//space complexity: O(n): capacity
//time complexity: O(1)
func (queue *Queue)Enqueue(data int) (error){
	//validate overflow
	rear:=((queue.Rear+1)%queue.Capacity)
	if (rear==queue.Front){
		return  errors.New("full queue exception:try after removing existing data")
	}
	queue.Elements[rear]=data
	queue.Rear=rear
	//handle first time addition to Q
	if queue.Front==-1{
		queue.Front=rear
	}
	return nil
}

//time complexity: t(n)=O(1)
func (queue *Queue)IsEmpty()bool{
	//validate underflow
	if queue.Front==-1{
		return true
	}
	return false
}

//time complexity: O(1)
func (queue *Queue)Dequeue() (int,error){
	if queue.IsEmpty(){
		return  0,errors.New("empty queue exception:try after adding data")
	}
	front:=queue.Front
	data:=queue.Elements[front]
	//handle when equal rear & front operations performed
	if queue.Front==queue.Rear{
		queue.Front=-1
		queue.Rear=-1
	}else{
		queue.Front=(queue.Front+1)%queue.Capacity
	}
	return data,nil
}

//t(n)=O(n)
func (queue *Queue) ReSize(additionalCapacity int) {
	queue.Mutex.Lock()
	//create new queue
	nElemens:=make([]int,queue.Capacity+additionalCapacity)
	//change existing capacity with additional one
	queue.Capacity=queue.Capacity+additionalCapacity
	//copy existing queue with same order of sequence
	for index,elem:=range queue.Elements{
		nElemens[index]=elem
	}
	queue.Elements=nElemens
	queue.Mutex.Unlock()
}

//t(n)=O(n)
func (queue *Queue) Delete(){
	queue.Elements=nil
}

//t(n)=O(n)
func (queue *Queue) Size()int{
	return (queue.Capacity-(queue.Rear+1 + queue.Front))%queue.Capacity
}

