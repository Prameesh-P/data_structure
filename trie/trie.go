package main

import "fmt"

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
		},
	}
}
func (t *Trie) Display() {
	t.displayHelper(t.root, []rune{})
}

func (t *Trie) displayHelper(node *TrieNode, prefix []rune) {
	if node.isEnd {
		fmt.Println(string(prefix))
	}
	for i, child := range node.children {
		if child != nil {
			t.displayHelper(child, append(prefix, 'a'+rune(i)))
		}
	}
}
func (t *Trie) Insert(word string) {
	node := t.root
	for _, r := range word {
		child, ok := node.children[r]
		if !ok {
			child = &TrieNode{
				children: make(map[rune]*TrieNode),
			}
			node.children[r] = child
		}
		node = child
	}
	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t.root
	for _, r := range word {
		child, ok := node.children[r]
		if !ok {
			return false
		}
		node = child
	}
	return node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t.root
	for _, r := range prefix {
		child, ok := node.children[r]
		if !ok {
			return false
		}
		node = child
	}
	return true
}

func main() {
	fmt.Println("Hello, 世界")
	t := NewTrie()
	t.Insert("aaa")
}
