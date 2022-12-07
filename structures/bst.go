package structures

import "golang.org/x/exp/constraints"

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
	case value < bst.Data && bst.Left != nil:
		return bst.Left.Search(value)
	case value > bst.Data && bst.Right != nil:
		return bst.Right.Search(value)
	}

	return true
}
