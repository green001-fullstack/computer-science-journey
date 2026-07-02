package transversal

import(
	"fmt"
)

type Node struct{
	Value int
	Left *Node
	Right *Node
}

func Inorder(node *Node){
	if node == nil{
		return
	}

	Inorder(node.Left)
	fmt.Println(node.Value)
	Inorder(node.Right)
}