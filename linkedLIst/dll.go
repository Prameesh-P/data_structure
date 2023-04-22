package main

import "fmt"

type Node struct {
    data int
    next *Node
    prev *Node
}

type DoublyLinkedList struct {
    head *Node
    tail *Node
}

func (dll *DoublyLinkedList) append(data int) {
    newNode := &Node{data: data}
    if dll.head == nil {
        dll.head = newNode
        dll.tail = newNode
    } else {
        dll.tail.next = newNode
        newNode.prev = dll.tail
        dll.tail = newNode
    }
}

func (dll *DoublyLinkedList) delete(data int) {
    if dll.head == nil {
        return
    }
    if dll.head.data == data {
        dll.head = dll.head.next
        if dll.head != nil {
            dll.head.prev = nil
        } else {
            dll.tail = nil
        }
        return
    }
    if dll.tail.data == data {
        dll.tail = dll.tail.prev
        dll.tail.next = nil
        return
    }
    currentNode := dll.head
    for currentNode != nil && currentNode.data != data {
        currentNode = currentNode.next
    }
    if currentNode != nil {
        currentNode.prev.next = currentNode.next
        currentNode.next.prev = currentNode.prev
    }
}

func (dll *DoublyLinkedList) printInOrder() {
    currentNode := dll.head
    for currentNode != nil {
        fmt.Printf("%d ", currentNode.data)
        currentNode = currentNode.next
    }
    fmt.Println()
}

func (dll *DoublyLinkedList) printInReverseOrder() {
    currentNode := dll.tail
    for currentNode != nil {
        fmt.Printf("%d ", currentNode.data)
        currentNode = currentNode.prev
    }
    fmt.Println()
}

func main() {
    dll := &DoublyLinkedList{}

    dll.append(1)
    dll.append(2)
    dll.append(3)
    dll.append(4)
    dll.append(5)

    dll.printInOrder() // Output: 1 2 3 4 5
    dll.printInReverseOrder() // Output: 5 4 3 2 1

    dll.delete(3)
    dll.delete(1)

    dll.printInOrder() // Output: 2 4 5
    dll.printInReverseOrder() // Output: 5 4 2
}