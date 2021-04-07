package queue

import (
	"sync"
)

type RemovalRequest struct {
	pos int
	id  string
}

//A Player queue works like a queue structure in an thread safe way.
type RemovalQueue struct {
	lock  sync.Mutex
	queue []RemovalRequest
}

func NewRemoval() *RemovalQueue {
	return &RemovalQueue{
		queue: make([]RemovalRequest, 0),
	}
}

//Enqueue enqueues a player that is added to the lobby when a position is free
func (q *RemovalQueue) Enqueue(player *RemovalRequest) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.queue = append(q.queue, *player)
}

//Dequeue dequeues one player from the waiting queue
func (q *RemovalQueue) Dequeue() *RemovalRequest {
	q.lock.Lock()
	defer q.lock.Unlock()
	if len(q.queue) > 0 {
		player := q.queue[0]
		q.queue = q.queue[1:]
		return &player
	}

	return nil
}

func (q *RemovalQueue) Length() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return len(q.queue)
}
