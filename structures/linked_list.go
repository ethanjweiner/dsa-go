package structures

import "fmt"

type LinkedList struct {
	Head   *Node
	Length int
}

type Node struct {
	Data any
	Next *Node
}

func (l *LinkedList) Search(value any) bool {
	currentNode := l.Head

	for currentNode != nil {
		if currentNode.Data == value {
			return true
		}
		currentNode = currentNode.Next
	}

	return false
}

func (l *LinkedList) Insert(value any) {
	l.Prepend(&Node{Data: value})
}

func (l *LinkedList) Prepend(n *Node) {
	n.Next = l.Head
	l.Head = n
	l.Length++
}

func (l *LinkedList) Print() {
	currentNode := l.Head

	for currentNode != nil {
		fmt.Println(currentNode.Data)
		currentNode = currentNode.Next
	}
}

// head -> head.next -> head.next.next
// 1 -> 2 -> 3 -> 4
func (l *LinkedList) DeleteWithValue(value any) {
	currentNode := l.Head

	if currentNode == nil {
		return
	}

	if currentNode.Data == value {
		l.Head = currentNode.Next
		l.Length--
		return
	}

	for currentNode.Next != nil {
		if currentNode.Next.Data == value {
			currentNode.Next = currentNode.Next.Next
			l.Length--
			return
		}
		currentNode = currentNode.Next
	}
}
