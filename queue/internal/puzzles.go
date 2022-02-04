package internal

//t(n)=O(n) ~n+n=2N
func ReverseQueue(queue *Queue) {
	//create a stack and push elements
	//var stack int;
	 for queue.IsEmpty(){
		//stack.push(queue.Dequeue())
	}
	//display the stack by pop
}

type TwoStackQueue struct{
	EnqueueStack []int
	DeQueueStack []int
}
func (stack *TwoStackQueue) EnQueue(){
	//stack.	EnqueueStack []int.Push
}
//t(n)_= O(1) single operation: ~O(1)+ O(n) once DequeStack is empty
func (stack *TwoStackQueue) DeQueue() {
	//stack.DeQueueStack.IsEmpty()
	//if Empty:copy from EnQueue till n-1
	//stack.EnQueueStack.Pop=stack.DeQueueStack.Push
	//else: if  stack.DeQueueStack.IsEmpty, throws error
	//stack.DeQueueStack.Pop
}

type StackByTwoQueue struct{
	Q1 Queue
	Q2 Queue
}

func (stack *StackByTwoQueue)IsEmpty(){

}

func (stack *StackByTwoQueue)IsFull(){

}

func (stack *StackByTwoQueue)Pop()int{
	//reverse one Q to another for extracting last element
	if stack.Q2.IsEmpty(){
		size:=stack.Q1.Size()
		index:=0
		//empty q1 to receive push ops
		for (index<(size-1)){
			elem,_:=stack.Q1.Dequeue()
			stack.Q1.Enqueue(elem)
			index++
		}
		lastElem,_:=stack.Q1.Dequeue()
		return lastElem
	}else{
		size:=stack.Q2.Size()
		index:=0
		for (index<(size-1)){
			elem,_:=stack.Q2.Dequeue()
			stack.Q1.Enqueue(elem)
		}
		lastElem,_:=stack.Q2.Dequeue()
		return lastElem
	}
}

func (stack *StackByTwoQueue)Push(data int){
	//check which queue is not empty
	if stack.Q1.IsEmpty(){
		stack.Q2.Enqueue(data)
	}else{
		stack.Q1.Enqueue(data)
	}
}

func SumOfSlidingWindow(window int,array []int)(int,[]int){
	//solved by priority queues
	return 0,[]int{}
}

func FindPairOfConsecutiveStack(stack []int){
	//create a queue
	//push each element to queue
	//Deque a pair form Queue
	//push back to stack again

}

func InterLeavingQueueHalf(){

}