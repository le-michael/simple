package simple

import (
	"reflect"
	"testing"
)

func TestStack_Empty(t *testing.T) {
	type fields struct {
		top *stackNode
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"Check new stack", fields{nil}, true},
		{"Check non empty stack", fields{&stackNode{4, nil}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := &Stack{
				top: tt.fields.top,
			}
			if got := stack.Empty(); got != tt.want {
				t.Errorf("Stack.Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Top(t *testing.T) {
	type fields struct {
		top *stackNode
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{"Check empty stack", fields{nil}, nil},
		{"Check non empty stack", fields{&stackNode{4, nil}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := &Stack{
				top: tt.fields.top,
			}
			if got := stack.Top(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stack.Top() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Push(t *testing.T) {

	stack := &Stack{nil}

	type fields struct {
		top *stackNode
	}

	tests := []struct {
		name string
		want interface{}
	}{
		{"Push to empty stack", 1},
		{"Push to non empty stack", 2},
		{"Push differnt type", "3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack.Push(tt.want)
			if tt.want != stack.top.value {
				t.Errorf("stack.top.value = %v, want %v", stack.top.value, tt.want)
			}
		})
	}
}

func TestStack_Pop(t *testing.T) {

	stack := &Stack{nil}
	stack.Push(1)
	stack.Push(2)

	type fields struct {
		top *stackNode
	}
	tests := []struct {
		name      string
		wantValue interface{}
		wantErr   bool
	}{
		{"Pop Normal", 1, false},
		{"Pop one item", nil, false},
		{"Pop one item", nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := stack.Pop(); (err != nil) != tt.wantErr || stack.Top() != tt.wantValue {
				t.Errorf("Stack.Pop() error = %v, wantErr %v, Stack.Top() = %v, wantValue=%v", err, tt.wantErr, stack.Top(), tt.wantValue)
			}
		})
	}
}

func TestNewStack(t *testing.T) {
	tests := []struct {
		name string
		want *Stack
	}{
		{"Init Stack", &Stack{nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStack(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStack() = %v, want %v", got, tt.want)
			}
		})
	}
}
