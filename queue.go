package simple

import "errors"

type (
	queueNode struct {
		value interface{}
		next  *queueNode
	}
	Queue struct {
		front, back *queueNode
	}
)

func (queue *Queue) Empty() bool {
	return queue.front == nil
}

func (queue *Queue) Front() interface{} {
	if queue.Empty() {
		return nil
	}
	return queue.front.value
}

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

func NewQueue() *Queue {
	return &Queue{nil, nil}
}
