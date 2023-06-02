package main

import "fmt"


type Node struct {
    Key   string
    Value interface{}
    Next  *Node
}

type HashTable struct {
    data []*Node
}

func NewHashTable(size int) *HashTable {
    return &HashTable{make([]*Node, size)}
}

func (ht *HashTable) hash(key string) int {
    hash := 0
    for i := 0; i < len(key); i++ {
        hash = (hash << 5) + hash + int(key[i])
        hash &= hash // keep the hash value positive
    }
    return hash % len(ht.data)
}

func (ht *HashTable) Get(key string) (interface{}, bool) {
    index := ht.hash(key)
    head := ht.data[index]
    for head != nil {
        if head.Key == key {
            return head.Value, true
        }
        head = head.Next
    }
    return nil, false
}

func (ht *HashTable) Put(key string, value interface{}) {
    index := ht.hash(key)
    head := ht.data[index]
    for head != nil {
        if head.Key == key {
            head.Value = value
            return
        }
        head = head.Next
    }
    newNode := &Node{Key: key, Value: value, Next: ht.data[index]}
    ht.data[index] = newNode
}

func (ht *HashTable) Remove(key string) {
    index := ht.hash(key)
    head := ht.data[index]
    var prev *Node
    for head != nil {
        if head.Key == key {
            if prev == nil {
                ht.data[index] = head.Next
            } else {
                prev.Next = head.Next
            }
            return
        }
        prev = head
        head = head.Next
    }
}
func main(){

	ht:=NewHashTable(10)
	ht.Put("a",1)
	ht.Put("b",3)
	h,_:=ht.Get("a")
	fmt.Println(h)	
}