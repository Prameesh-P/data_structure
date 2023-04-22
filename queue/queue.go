package main

import (
	. "fmt"
	"log"
)

type Node struct {
	data interface{}
	next *Node
}
type Queue struct {
	head *Node
	tail *Node
}

func (q *Queue) EnQueue(data interface{}) {

	newNode := &Node{data: data, next: nil}
	if q.tail == nil {
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}

}
func (q *Queue)ReverseDispaly(){

	var val []interface{}
	curr:=q.head
	if curr == nil{
		log.Fatalf("Stack Underflow ")
	}

	for curr !=nil{
		val = append(val, curr.data)
		curr=curr.next
	}
	Print("[")
	for i:=len(val)-1;i>=0;i--{

		Print(val[i])

	}
	Print("]")
	Println()

}
func (q *Queue) Display() {
    if q.head == nil {
        Println("Queue is empty")
        return
    }
    Print("Queue: ")
    current := q.head
    for current != nil {
        Printf("%v ", current.data)
        current = current.next
    }
    Println()
}
func (q *Queue) DeQueue() interface{} {

	if q.head == nil {
		return nil
	}
	val := q.head.data
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	return val
}

func main() {
	q:=&Queue{}
	q.EnQueue(5)
	q.EnQueue(6)
	q.EnQueue(7)
	q.Display()
	q.ReverseDispaly()
	q.DeQueue()
	q.Display()
	q.DeQueue()
	q.Display()
	q.DeQueue()
	q.Display()
}
