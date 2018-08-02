package simple

import (
	"reflect"
	"testing"
)

func TestNewQueue(t *testing.T) {
	tests := []struct {
		name string
		want *Queue
	}{
		{"Init Queue", &Queue{nil, nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewQueue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Pop(t *testing.T) {

	n0 := &queueNode{1, nil}
	n1 := &queueNode{2, nil}
	n2 := &queueNode{3, nil}
	n0.next = n1
	n1.next = n2

	type fields struct {
		front *queueNode
		back  *queueNode
	}
	tests := []struct {
		name      string
		fields    fields
		wantErr   bool
		wantFront *queueNode
		wantBack  *queueNode
	}{
		{"Pop empty Queue", fields{nil, nil}, true, nil, nil},
		{"Pop normal", fields{n0, n2}, false, n1, n2},
		{"Pop two items in queue", fields{n1, n2}, false, n2, n2},
		{"Pop one item in queue", fields{n2, n2}, false, nil, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := &Queue{
				front: tt.fields.front,
				back:  tt.fields.back,
			}
			if err := queue.Pop(); (err != nil) != tt.wantErr || queue.front != tt.wantFront || queue.back != tt.wantBack {
				t.Errorf("Queue.Pop() error = %v, wantErr %v, front = %v, wantFront = %v, back = %v, wantBack = %v", err, tt.wantErr, queue.front, tt.wantFront, queue.back, tt.wantBack)
			}
		})
	}
}

func TestQueue_Empty(t *testing.T) {
	type fields struct {
		front *queueNode
		back  *queueNode
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"Queue on init", fields{nil, nil}, true},
		{"Queue has item", fields{&queueNode{0, nil}, &queueNode{0, nil}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := &Queue{
				front: tt.fields.front,
				back:  tt.fields.back,
			}
			if got := queue.Empty(); got != tt.want {
				t.Errorf("Queue.Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Front(t *testing.T) {
	type fields struct {
		front *queueNode
		back  *queueNode
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{"Queue on init", fields{nil, nil}, nil},
		{"Queue has item", fields{&queueNode{0, nil}, &queueNode{0, nil}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := &Queue{
				front: tt.fields.front,
				back:  tt.fields.back,
			}
			if got := queue.Front(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Queue.Front() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Push(t *testing.T) {
	type fields struct {
		front *queueNode
		back  *queueNode
	}
	type args struct {
		value interface{}
	}
	tests := []struct {
		name      string
		args      interface{}
		wantFront interface{}
		wantBack  interface{}
	}{
		{"Push to empty queue", 1, 1, 1},
		{"Push to queue with one item", 2, 1, 2},
		{"Push to queue with differnt type", "3", 1, "3"},
	}
	queue := &Queue{nil, nil}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue.Push(tt.args)
			if tt.wantBack != queue.back.value || tt.wantFront != queue.front.value {
				t.Errorf("front.value = %v, wantFront = %v, back.value = %v, wantBack = %v", queue.front.value, tt.wantFront, queue.back.value, tt.wantBack)
			}
		})
	}
}
