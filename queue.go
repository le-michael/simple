package simple

import "errors"

type (
	queueNode struct {
		value interface{}
		next  *queueNode
	}
	//Queue struct
	Queue struct {
		front, back *queueNode
	}
)

//Empty returns if the queue is empty as a boolean
func (queue *Queue) Empty() bool {
	return queue.front == nil
}

//Front returns the value of the front of the queue
func (queue *Queue) Front() interface{} {
	if queue.Empty() {
		return nil
	}
	return queue.front.value
}

//Push a value to the back of the queue
func (queue *Queue) Push(value interface{}) {
	node := &queueNode{value, nil}
	if queue.Empty() {
		queue.front = node
		queue.back = node
		return
	}

	queue.back.next = node
	queue.back = node
}

//Pop removes a value from the front of the queue
func (queue *Queue) Pop() error {
	if queue.Empty() {
		return errors.New("queue: cannot pop from empty queue")
	}

	queue.front = queue.front.next
	if queue.Empty() {
		queue.back = nil
	}

	return nil
}

//NewQueue creates a new queue object
func NewQueue() *Queue {
	return &Queue{nil, nil}
}
