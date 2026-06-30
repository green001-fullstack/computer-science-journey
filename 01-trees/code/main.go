package main
import (
	"tree/implementations"
)



func main(){

	root := &implementations.Node{
		Value : 15,
	}

	nodeTen := &implementations.Node{ Value : 10}
	nodeTwentyFive := &implementations.Node{ Value : 25}
	nodeEight := &implementations.Node{ Value : 8}

	root.Left = nodeTen
	root.Right = nodeTwentyFive
	nodeTen.Left = nodeEight
	implementations.Inorder(root)
}