package main

import "fmt"

type Node struct {
	Value string
	Next  *Node
}

func NewNode(value string) *Node {
	return &Node{Value: value}
}

func (n *Node) Insert(node *Node) {
	node.Next = n.Next
	n.Next = node
}

type LinkedList struct {
	Head *Node
}

func NewLinkedList() *LinkedList {
	node := &Node{}
	node.Next = node
	return &LinkedList{
		Head: node,
	}
}

func (l *LinkedList) InsertHead(node *Node) {
	l.Head.Insert(node)
}

func (l *LinkedList) Print() {
	node := l.Head.Next
	for ; node != l.Head; node = node.Next {
		fmt.Printf("-> %v ", node.Value)
	}
	fmt.Println("")
}

func main() {
	list := NewLinkedList()
	list.InsertHead(NewNode("foo"))
	list.InsertHead(NewNode("bar"))
	list.Head.Next.Insert(NewNode("baz"))
	list.Print()
}
