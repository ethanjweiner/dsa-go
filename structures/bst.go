package structures

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type BST[T constraints.Ordered] struct {
	Data  T
	Left  *BST[T]
	Right *BST[T]
}

func (bst *BST[T]) Insert(value T) {
	switch {
	case value < bst.Data && bst.Left == nil:
		bst.Left = &BST[T]{Data: value}
	case value < bst.Data && bst.Left != nil:
		bst.Left.Insert(value)
	case value > bst.Data && bst.Right == nil:
		bst.Right = &BST[T]{Data: value}
	case value > bst.Data && bst.Right != nil:
		bst.Right.Insert(value)
	}
}

func (bst *BST[T]) Search(value T) bool {
	switch {
	case bst == nil:
		return false
	case value < bst.Data:
		return bst.Left.Search(value)
	case value > bst.Data:
		return bst.Right.Search(value)
	}

	return true
}

// Returns updated tree as expected
func (bst *BST[T]) Delete(value T) *BST[T] {
	switch {
	case bst == nil:
		return bst
	case value < bst.Data && bst.Left != nil:
		bst.Left = bst.Left.Delete(value)
		return bst
	case value > bst.Data && bst.Right != nil:
		bst.Right = bst.Right.Delete(value)
		return bst
	}

	// We have found the value to delete

	// If one child, make that child the successor
	if bst.Left == nil {
		return bst.Right
	}

	if bst.Right == nil {
		return bst.Left
	}

	// Transform the deleted node into the successor node
	// If the successor node has right children, `lift` will set it's parent's
	// `Left` to those children
	bst.Right = bst.Right.lift(bst)
	return bst
}

func (bst *BST[T]) lift(nodeToDelete *BST[T]) *BST[T] {
	if bst.Left != nil {
		bst.Left = bst.Left.lift(nodeToDelete)
		return bst
	}

	// Reached the successor
	nodeToDelete.Data = bst.Data
	return bst.Right
}

func (bst *BST[T]) TraverseInOrder(f func(T)) {
	if bst == nil {
		return
	}

	bst.Left.TraverseInOrder(f)
	f(bst.Data)
	bst.Right.TraverseInOrder(f)
}

func (bst *BST[T]) PrintInOrder() {
	bst.TraverseInOrder(func(data T) {
		fmt.Printf("%v ", data)
	})
}
