package lesson15bfs

func (q *Queue) Dequeue() *Node {
	if len(q.items) == 0{
		return nil
	}

	current := q.items[0]
	q.items = q.items[1:]
	return current
}