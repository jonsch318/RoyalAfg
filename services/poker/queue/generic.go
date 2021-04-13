package queue

import (
	"sync"
)

type IQueue interface {
	Enqueue(interface{})
	Dequeue() interface{}
	Length() int
}

//A Player queue works like a queue structure in an thread safe way.
type Queue struct {
	lock  sync.Mutex
	queue []interface{}
}

func New() *Queue {
	return &Queue{
		queue: make([]interface{}, 0),
	}
}

//Enqueue enqueues a player that is added to the lobby when a position is free
func (q *Queue) Enqueue(player interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.queue = append(q.queue, player)
}

//Dequeue dequeues one player from the waiting queue
func (q *Queue) Dequeue() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()
	if len(q.queue) > 0 {
		player := q.queue[0]
		q.queue = q.queue[1:]
		return player
	}

	return nil
}

func (q *Queue) Length() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return len(q.queue)
}
