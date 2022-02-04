package internal

import "errors"
//keep track of only rear node and front node
type LinkedListQueue struct {
	Rear *ListNode
	Front *ListNode
}

//nodes in the linked list
type ListNode struct{
	Next *ListNode
	Data int
}

func NewLinkedListQueue()LinkedListQueue{
	queue:= LinkedListQueue{
		Rear:nil,
		Front: nil,
	}
	return queue
}

func (lQueue *LinkedListQueue)IsEmpty()bool{
	return lQueue.Front==nil
}

func (lQueue *LinkedListQueue)IsFull()bool{
	//create new node till ram memory is full
	return false
}

func (lQueue *LinkedListQueue) EnQueue(data int)error{
	if lQueue.IsFull(){
		return errors.New("queue is full kindly remove some of the existing and add later")
	}
	//create node first
	nNode:=&ListNode{
		Data: data,
		Next: nil,
	}

	//handle front is empty
	if lQueue.Front==nil{
		lQueue.Front=lQueue.Rear
	}
	//first time rear node is null
	if lQueue.Rear!=nil{
		lQueue.Rear.Next=nNode
	}
	//assign node to current rear
	lQueue.Rear=nNode
	return nil
}

func (lQueue *LinkedListQueue)DeQueue()(int,error){
	if lQueue.IsEmpty(){
		return 0,errors.New("queue is empty kindly add data and try later")
	}
	//fetch data in current node
	front:=	lQueue.Front
	data:=front.Data
	lQueue.Front=lQueue.Front.Next
	front=nil //free up front
	return data,nil
}