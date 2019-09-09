package main

import "github.com/gw123/algorithm/linkedList"

func main() {
	list := linkedList.NewLinkedList()
	list.Append("1")
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append("123123123")
	list.Print()
}
