package main

type circularQueue struct{
	nodes []*node
	size,head,last int
}

type node struct {
	data int
}

func (queue *circularQueue) new(){
	queue.nodes=make([]*node,10)
	queue.head=0
	queue.last=0
	queue.size=10
}

func unUsed(){
	//check head & last are same means no elements added
}

func isComplete(){
	//if head==(last+1)%size
}

func (queue *circularQueue)add(data int){
	isComplete()//return error
	queue.nodes[queue.last]=&node{data}
	queue.last=(queue.last+1)%queue.size
}

func (queue *circularQueue)moveHead()*node{
	head:=queue.nodes[queue.head]
	queue.head=(queue.head+1)%queue.size
	return head
}