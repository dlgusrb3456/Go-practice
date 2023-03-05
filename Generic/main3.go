package main

import (
	"fmt"
)

type Node[T any] struct {
	val  T
	next *Node[T]
}

func NewNode[T any](v T) *Node[T] {
	return &Node[T]{v, nil} // or &Node[T]{val:v}
}

func (n *Node[T]) Push(v T) *Node[T] { // 단, 메소드에는 Generic Type을 새로운 인자로 선언해서 넣을 수 없음
	node := NewNode(v) // method must have no type parameters
	n.next = node
	return node
}

func main() {
	fmt.Println("Generic #3, Generic Type")

	{
		node1 := NewNode(1)     // Type : *Node[int]
		node2 := NewNode("Two") // Type : *Node[string]
		node3 := NewNode(3)
		node1.next = node3 // node2 Type is not *Node[int]

		node3.Push(4).Push(5).Push(6)
		node2.Push("Three").Push("Four")

		for node1 != nil {
			fmt.Println(node1.val, " ")
			node1 = node1.next
		}

		for node2 != nil {
			fmt.Println(node2.val)
			node2 = node2.next
		}
	}

	// Generic 좋긴 한데 코드 가독성이 좀... (공감함)
	// Working code => Refactoring (이 단계에서 추가해도 늦지 않음)
}
