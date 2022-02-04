package main

import "fmt"

type LinkedList struct{
	head *Node
}

type Node struct{
	data int
	alphabet rune
	next *Node
}

func (list *LinkedList)Create(){
	head:=&Node{
		alphabet:'d',
	}
	var node=head
	for _,character:=range "inil"{
		fmt.Println("Character",character)
		node.next=	&Node{
			alphabet: character,
		}
		node=node.next
	}
	list.head=head
}

func (list *LinkedList)Traversal(){
	node:=list.head
	for node!=nil{
		fmt.Println("Node Value:-",node.alphabet)
		node=node.next
	}

}
func (list *LinkedList)Reversal(){
	node:=list.head
	recurReversal(node)
}

func recurReversal(node *Node){
	if node!=nil{
		recurReversal(node.next)
		fmt.Println("node value reversed:",node.alphabet)
	}
	return
}
func main(){
	list:=&LinkedList{}
	list.Create()
	list.Traversal()
	list.Reversal()
}

