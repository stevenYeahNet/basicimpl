package basiciml

import (
	"testing"
)

func TestQueue(t *testing.T)  {
	var q Queue = *NewQueue()
	var data int = 1
	q.Push(data)
	data = 2
	q.Push(data)
	data = 3
	q.Push(data)
	ff,_ := q.Front()
	if ff != 1{
		t.Error("error data")
	}
	q.Pop()
	ff,_ = q.Front()
	if ff != 2{
		t.Error("error data")
	}
	q.Pop()
	ff,_ = q.Front()
	if ff != 3{
		t.Error("error data")
	}
}
