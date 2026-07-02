package main
import (
	"tree/bfs"
)



func main(){

	root := &bfs.Node{
		Value : 15,
	}

	nodeTen := &bfs.Node{ Value : 10}
	nodeTwentyFive := &bfs.Node{ Value : 25}
	nodeEight := &bfs.Node{ Value : 8}

	root.Left = nodeTen
	root.Right = nodeTwentyFive
	nodeTen.Left = nodeEight
	bfs.BFS(root)
}