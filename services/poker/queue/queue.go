package queue

import (
	"sync"

	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"
)

//A Player queue works like a queue structure in an thread safe way.
type PlayerQueue struct {
	lock sync.Mutex
	queue []*models.Player
}

func New() *PlayerQueue {
	return &PlayerQueue{
		queue: make([]*models.Player, 0),
	}
}

//Enqueue enqueues a player that is added to the lobby when a position is free
func (q *PlayerQueue) Enqueue(player *models.Player) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.queue = append(q.queue, player)
}

//Dequeue dequeues one player from the waiting queue
func (q *PlayerQueue) Dequeue() *models.Player {
	q.lock.Lock()
	defer q.lock.Unlock()
	if len(q.queue) > 0 {
		player := q.queue[0]
		q.queue = q.queue[1:]
		return player
	}

	return nil
}

func (q *PlayerQueue) Length() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return len(q.queue)
}