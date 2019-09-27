package basiciml

import (
	"errors"
)

type Queue struct {
	head *Node
}

func NewQueue() *Queue {
	s := &Queue{nil}
	return s
}
func (q *Queue) Push(data interface{}) {
	node := &Node{Node:data,next:nil}
	if q.head == nil {
		q.head = node
		return
	}
	tmp := q.head
	for{
		if tmp.next == nil{
			break
		}
		tmp = tmp.next
	}
	tmp.next = node
}
func (q *Queue) Front() (interface{},error){
	if q.Size() == 0{
		return nil,errors.New("queue is empty")
	}
	tmp := q.head
	var res interface{} = tmp.Node
	return res,nil
}
func (q *Queue) Size() int {
	_size := 0
	tmp := q.head
	for {
		if tmp == nil{
			break
		}
		_size++
		tmp = tmp.next
	}
	return _size
}

func (q *Queue) Pop() error{
	if q.Size() == 0{
		return errors.New("Can't pop an empty queue")
	}
	tmp := q.head
	tmp = tmp.next
	q.head = tmp
	return nil
}

func (q *Queue) Empty() bool {
	return q.Size() == 0
}