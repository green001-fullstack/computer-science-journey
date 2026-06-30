package main

type Node struct{
	Value int
	Left *Node
	Right *Node
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
}