package main

import "fmt"

type Node struct {
    data int
    next *Node
}

type LinkedList struct {
    head *Node
}

func (ll *LinkedList) addNode(data int) {
    newNode := &Node{data: data}
    if ll.head == nil {
        ll.head = newNode
    } else {
        curr := ll.head
        for curr.next != nil {
            curr = curr.next
        }
        curr.next = newNode
    }
}

func (ll *LinkedList) removeNode(data int) {
    if ll.head == nil {
        return
    }

    if ll.head.data == data {
        ll.head = ll.head.next
        return
    }

    curr := ll.head
    for curr.next != nil {
        if curr.next.data == data {
            curr.next = curr.next.next
            return
        }
        curr = curr.next
    }
}

func (ll *LinkedList) display() {
    curr := ll.head
    for curr != nil {
        fmt.Printf("%d -> ", curr.data)
        curr = curr.next
    }
    fmt.Println("nil")
}
func (ll *LinkedList) printInReverseOrder() {
    printInReverseOrderHelper(ll.head)
    fmt.Println()
}

func printInReverseOrderHelper(node *Node) {
    if node == nil {
        return
    }
    printInReverseOrderHelper(node.next)
    fmt.Printf("%d ", node.data)
}

func main() {
    ll := &LinkedList{}
    ll.addNode(1)
    ll.addNode(2)
    ll.addNode(3)
    ll.addNode(4)
	ll.printInReverseOrder()
    ll.display()

    ll.removeNode(3)
    ll.display()

    ll.removeNode(1)
    ll.display()
}