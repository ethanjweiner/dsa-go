package practice_problems

import (
	. "github.com/ethanjweiner/dsa-go/structures"
	"github.com/ethanjweiner/dsa-go/utils"
)

// DFS normal
func findDepth[T any](bt *BT[T]) int {
	switch {
	case isLeaf(bt):
		return 1
	case bt.Left == nil:
		return 1 + findDepth(bt.Right)
	case bt.Right == nil:
		return 1 + findDepth(bt.Left)
	}

	leftDepth := findDepth(bt.Left)
	rightDepth := findDepth(bt.Right)

	return 1 + utils.Min(leftDepth, rightDepth)
}

// DFS accum
func findDepth[T any](bt *BT[T]) int {
	return findRemainingDepth(bt, 0)
}

func findRemainingDepth[T any](bt *BT[T], depthSoFar int) int {
	switch {
	case isLeaf(bt):
		return 1 + depthSoFar
	case bt.Left == nil:
		return findRemainingDepth(bt.Right, 1+depthSoFar)
	case bt.Right == nil:
		return findRemainingDepth(bt.Left, 1+depthSoFar)
	}

	leftDepth := findRemainingDepth(bt.Left, 1+depthSoFar)
	rightDepth := findRemainingDepth(bt.Right, 1+depthSoFar)

	return utils.Min(leftDepth, rightDepth)
}

// BFS
func findDepth[T any](bt *BT[T]) int {
	depthSoFar := 1

	queue := &Queue[*BT[T]]{}
	queue.Enqueue(bt)

	for {
		for i := 0; i < queue.Length(); i++ {
			dequeued := queue.Dequeue()

			switch {
			case isLeaf(dequeued):
				return depthSoFar
			case dequeued.Left != nil:
				queue.Enqueue(dequeued.Left)
			case dequeued.Right != nil:
				queue.Enqueue(dequeued.Right)
			}
		}

		depthSoFar++
	}

	// Go doesn't complain?
	// Alternative: Return min tree depth
}

func isLeaf[T any](bt *BT[T]) bool {
	return bt.Left == nil && bt.Right == nil
}
