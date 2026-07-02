package lesson15bfs

import "fmt"

// type Queue struct{
// 	items []*Node
// }

// func(q *Queue) Enqueue(node *Node){
// 	q.items = append(q.items, node)
// }

// func (q *Queue) Dequeue() *Node {
// 	if len(q.items) == 0{
// 		return nil
// 	}

// 	current := q.items[0]
// 	q.items = q.items[1:]
// 	return current
// }

// func(q *Queue) IsEmpty()bool{
// 	return len(q.items) == 0
// }

func BFS(root *Node){
	if root == nil{
		return
	}

	q := Queue{}
	q.Enqueue(root)

	for !q.isEmpty(){
		current := q.Dequeue()

		fmt.Println(current.Value)
		if current.Left != nil{
			q.Enqueue(current.Left)
		}
		if current.Right != nil{
			q.Enqueue(current.Right)
		}
	}


}