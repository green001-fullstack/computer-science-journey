package main
import (
	"fmt"
)

type Node struct{
	Value int
	Left *Node
	Right *Node
}

func Preorder(node *Node){
	if node == nil{
		return
	}

	fmt.Println(node.Value)
	Preorder(node.Left)
	Preorder(node.Right)
}

func main(){

	root := &Node{
		Value : 15,
	}

	nodeTen := &Node{ Value : 10}
	nodeTwentyFive := &Node{ Value : 25}
	nodeEight := &Node{ Value : 8}

	root.Left = nodeTen
	root.Right = nodeTwentyFive
	nodeTen.Left = nodeEight
	Preorder(root)
}