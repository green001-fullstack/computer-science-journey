package bfs

func(q *Queue) isEmpty()bool{
	return len(q.items) == 0
}