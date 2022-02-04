package doubly_linked_list

type doublyLinkedList struct {
	head *node
}

type node struct {
	next *node
	prev *node
	data int
}

func (list *doublyLinkedList)Create(){
	//head node separate implementation
	//loop through and insert data
}

func (list *doublyLinkedList)pushBack(data int){

}