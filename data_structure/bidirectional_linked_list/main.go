package main

import "fmt"

type Node struct {
	Value string
	Next  *Node
	Prev  *Node
}

func NewNode(value string) *Node {
	return &Node{Value: value}
}

func (n *Node) Insert(insertion *Node) {
	// 「n -> n.Next」を「n -> insertion -> n.Next」にする
	insertion.Next = n.Next
	n.Next.Prev = insertion
	n.Next = insertion
	insertion.Prev = n
}

func (n *Node) Erase() {
	// 「n.Prev -> n -> n.Next」を「n.Prev -> n.Next」にする
	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev
}

type BidirectionalLinkedList struct {
	Head *Node
}

func NewBidirectionalLinkedList() *BidirectionalLinkedList {
	node := &Node{}
	node.Next = node
	node.Prev = node
	return &BidirectionalLinkedList{
		Head: node,
	}
}

func (l *BidirectionalLinkedList) Erase(node *Node) {
	if node == l.Head {
		return
	}
	node.Erase()
}

func (l *BidirectionalLinkedList) InsertHead(node *Node) {
	l.Head.Insert(node)
}

func (l *BidirectionalLinkedList) Print() {
	node := l.Head.Next
	for ; node != l.Head; node = node.Next {
		fmt.Printf("-> %v ", node.Value)
	}
	fmt.Println("")
}

func main() {
	list := NewBidirectionalLinkedList()
	baz := NewNode("baz")

	list.InsertHead(NewNode("foo"))
	list.InsertHead(baz)
	list.InsertHead(NewNode("bar"))

	fmt.Printf("inserted: ")
	list.Print()

	list.Erase(baz)
	fmt.Printf("erased: ")
	list.Print()
}
