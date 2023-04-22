package main

import (
	. "fmt"
	"log"
)

type Node struct{

	data int
	next *Node

}

type Stack struct{

	top *Node

}

func (s *Stack)Push(data int){

	newNode:=&Node{data: data,next: s.top}
	s.top=newNode

}

func (s *Stack)ReverseDispaly(){

	var val []int
	curr:=s.top
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
func (s *Stack) Pop()int{

	if s.top == nil{
		log.Fatalf("Stack Underflow ")
	}
	value:=s.top.data
	s.top=s.top.next
	return value
}

func (s *Stack) Display(){

	curr:=s.top

	if curr == nil{
		log.Fatalf("Stack Underflow ")
	}

	Print("[")
	for curr !=nil{
		Print(curr.data)
		curr=curr.next
	}
	Print("]")
	Println()
}
func (s *Stack)Peek()int{
	return s.top.data
}


func main(){
	Println("hyy")
	s:=&Stack{}
	s.Push(8)
	s.Push(3)
	s.Push(6)
	s.Display()
	s.ReverseDispaly()
	Println(s.Peek())
	Println(s.Pop())
	s.Display()
}