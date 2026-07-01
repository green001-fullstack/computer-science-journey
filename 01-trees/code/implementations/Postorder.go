package implementations

import "fmt"

func PostOrder(node *Node) {
	if node == nil {
		return
	}

	PostOrder(node.Left)
	PostOrder(node.Right)
	fmt.Println(node.Value)
}