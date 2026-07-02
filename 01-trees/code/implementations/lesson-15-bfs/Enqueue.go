package lesson15bfs

func(q *Queue) Enqueue(node *Node){
	q.items = append(q.items, node)
}