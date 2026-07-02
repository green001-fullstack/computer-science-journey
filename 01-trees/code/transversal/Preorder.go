package transversal
import "fmt"

func Preorder(node *Node){
	if node == nil{
		return
	}

	fmt.Println(node.Value)
	Preorder(node.Left)
	Preorder(node.Right)
}