package linkedList

import "fmt"

type Node struct {
	Data interface{}
	Next *Node
	Pre  *Node
}

type LinkedList struct {
	Head *Node
	cur  *Node
}

func NewLinkedList() *LinkedList {
	head := &Node{}
	head.Next = head
	head.Pre = head
	linkedList := &LinkedList{}
	linkedList.Head = head
	linkedList.cur = head
	return linkedList
}

func (this *LinkedList) Append(data interface{}) error {
	node := &Node{Data: data}
	tempNode := this.Head.Next
	this.Head.Next = node
	node.Next = tempNode
	node.Pre = this.Head
	tempNode.Pre = node
	return nil
}

func(this *LinkedList)Next()  {
	this.cur = this.cur.Next
}

func (this *LinkedList)Pre()  {
	this.cur = this.cur.Pre
}

func (this *LinkedList)RemoveCurrnt()  {
	this.cur.Next.Pre = this.cur.Pre
	this.cur.Pre.Next = this.cur.Next
}

func (this *LinkedList) NextElem() interface{} {
	return this.cur.Next.Data
}

func (this *LinkedList) CurrentElem() interface{}{
	return this.cur.Data
}

func (this *LinkedList) PreElem() interface{}  {
	return this.cur.Pre.Data
}

func (this *LinkedList) Print() {
	tempNode := this.Head.Next
	for tempNode.Next != this.Head {
		fmt.Println(tempNode.Data)
		tempNode = tempNode.Next
	}
	fmt.Println(tempNode.Data)
}
