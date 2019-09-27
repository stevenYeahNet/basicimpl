package basiciml

import (
	"errors"
)

type Node struct {
	Node interface{}
	next *Node
}

type Stack struct {
	head *Node
}

func NewStack() *Stack {
	s := &Stack{nil}
	return s
}

func (s *Stack) Push(data interface{}) {
	node := &Node{Node:data,next:s.head}
	s.head = node
}

func (s *Stack) Pop() (interface{},bool) {
	n := s.head
	if( n == nil ){
		return nil,false
	}
	s.head = s.head.next
	return n.Node,true
}

//another imp
type _Stack []interface{}

func (s *_Stack) Push(data interface{}) bool {
	*s = append(*s,data)
	return true
}

func (s *_Stack) Pop() error {
	length := len(*s)
	if length == 0{
		return errors.New("Can't pop an empty stack")
	}
	*s = (*s)[0:length - 1]
	return nil
}

func (s *_Stack) Top() (interface{},error) {
	length := len(*s)
	if length == 0{
		return  nil,errors.New("stack is empty()!")
	}
	data := (*s)[length - 1]
	return data,nil
}

func (s *_Stack) Empty() bool {
	if( len(*s) == 0 ){
		return true
	}else {
		return false
	}
}