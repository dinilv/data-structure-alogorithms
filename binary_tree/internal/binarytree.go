package main

import (
	"fmt"
)

type BinaryTreeNode struct {
	  data int
	 leftNode *BinaryTreeNode
	 rightNode *BinaryTreeNode
}

type BinarySearchTree struct{
	root *BinaryTreeNode
}
func (tree *BinarySearchTree)InsertElement(data int){
	newNode:=&BinaryTreeNode{
		data:data,
	}
	if tree.root==nil{
		tree.root=newNode
		return
	}
	tree.InsertNode(tree.root,newNode)
}

func (tree *BinarySearchTree)InsertNode(current,newNode *BinaryTreeNode){
	//go left tree traversal
	if current.data>newNode.data{
		fmt.Println("Going left")
		if current.leftNode==nil{
			current.leftNode=newNode
		}else{
			tree.InsertNode(current.leftNode,newNode )
		}
	}
	if current.data<newNode.data{
		fmt.Println("Going right")
		if current.rightNode==nil{
			current.rightNode=newNode
		}else{
			tree.InsertNode(current.rightNode,newNode )
		}
	}
}

func (tree * BinarySearchTree)MinNode(){
	tree.recurMinNode(tree.root)
}
func (tree * BinarySearchTree)recurMinNode(currentNode *BinaryTreeNode){
	if currentNode.leftNode==nil{
		fmt.Println("Found min node",currentNode.data)
	}else{
		tree.recurMinNode(currentNode.leftNode)
	}
}
func (tree * BinarySearchTree)PreOrderTreeTraversal(){
	if tree.root==nil{
		fmt.Println("Binary Tree is empty.")
		return
	}
	tree.recurPreOrderTreeTraversal(tree.root)
}

func (tree * BinarySearchTree)RemoveNode(data int){
	// handle multiple conditions
}
func (tree *BinarySearchTree)recurPreOrderTreeTraversal(currentNode *BinaryTreeNode){
	if currentNode!=nil{
		fmt.Println("Pre Order Node:",currentNode.data)
		tree.recurInOrderTreeTraversal(currentNode.leftNode)
		tree.recurInOrderTreeTraversal(currentNode.rightNode)
	}
}
func (tree * BinarySearchTree)PostOrderTreeTraversal(){

}
func (tree *BinarySearchTree)InOrderTreeTraversal(){
	if tree.root==nil{
		fmt.Println("Binary Tree is empty.")
		return
	}
	tree.recurInOrderTreeTraversal(tree.root)
}
func (tree *BinarySearchTree)recurInOrderTreeTraversal(currentNode *BinaryTreeNode){
	if currentNode!=nil{
		tree.recurInOrderTreeTraversal(currentNode.leftNode)
		fmt.Println("In Order Node:",currentNode.data)
		tree.recurInOrderTreeTraversal(currentNode.rightNode)
	}
	return
}

func main()  {
	//create tree
	tree:=BinarySearchTree{}
	tree.InsertElement(3)
	tree.InsertElement(2)
	tree.InsertElement(4)
	tree.InsertElement(5)
	tree.InsertElement(6)
	tree.InsertElement(1)
	tree.MinNode()
}

