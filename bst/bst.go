package main

import (
    "fmt"
)

type Node struct {
    key   int
    left  *Node
    right *Node
}
type BST struct {
    root *Node
}
func NewNode(key int) *Node {
    return &Node{
        key:   key,
        left:  nil,
        right: nil,
    }
}



func NewBST() *BST {
    return &BST{
        root: nil,
    }
}

func (bst *BST) Insert(key int) {
    if bst.root == nil {
        bst.root = NewNode(key)
    } else {
        bst.root.insert(key)
    }
}

func (n *Node) insert(key int) {
    if key < n.key {
        if n.left == nil {
            n.left = NewNode(key)
        } else {
            n.left.insert(key)
        }
    } else if key > n.key {
        if n.right == nil {
            n.right = NewNode(key)
        } else {
            n.right.insert(key)
        }
    }
}

func (bst *BST) Search(key int) bool {
    return bst.root.search(key)
}

func (n *Node) search(key int) bool {
    if n == nil {
        return false
    } else if key == n.key {
        return true
    } else if key < n.key {
        return n.left.search(key)
    } else {
        return n.right.search(key)
    }
}

func main() {
    bst := NewBST()
    bst.Insert(10)
    bst.Insert(5)
    bst.Insert(15)
    bst.Insert(3)
    bst.Insert(7)
    bst.Insert(13)
    bst.Insert(17)
    fmt.Println(bst.Search(5)) // Output: true
    fmt.Println(bst.Search(12)) // Output: false
}
