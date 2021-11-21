package queues

import (
	"errors"
	"fmt"
)

type queue struct {
	items []int
}

func (q *queue) enqueue(newItem int) {
	q.items = append(q.items, newItem)
}

func (q *queue) deqeue() (int, error) {
	if len(q.items) > 0 {
		removed := q.items[0]
		q.items = q.items[1:]
		return removed, nil
	} else {
		return 0, errors.New("empty queue")
	}
}

func (q queue) length() int {
	return len(q.items)
}

func InitQueueExample() {
	myQueue := queue{}
	fmt.Println(myQueue)
	myQueue.enqueue(1)
	myQueue.enqueue(2)
	myQueue.enqueue(3)
	fmt.Println(myQueue)
	fmt.Println("Queue length", myQueue.length())
	removed1, _ := myQueue.deqeue()
	removed2, _ := myQueue.deqeue()
	fmt.Println("Removed: ", removed1, removed2)
	fmt.Println(myQueue)
	myQueue.deqeue()
	removedErr, err := myQueue.deqeue()
	fmt.Println(removedErr, err)
}
