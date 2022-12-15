package practice_problems

import (
	. "github.com/ethanjweiner/dsa-go/structures"
	"golang.org/x/exp/constraints"
)

func traverse[T constraints.Ordered](bt *BT[T]) [][]T {
	queue := &Queue[*BT[T]]{}
	queue.Enqueue(bt)
	levels := [][]T{}

	for !queue.IsEmpty() {
		newQueue := &Queue[*BT[T]]{}
		currentLevel := []T{}

		for !queue.IsEmpty() {
			dequeued := queue.Dequeue()
			currentLevel = append(currentLevel, dequeued.Data)
			if dequeued.Left != nil {
				newQueue.Enqueue(dequeued.Left)
			}

			if dequeued.Right != nil {
				newQueue.Enqueue(dequeued.Right)
			}
		}

		levels = append(levels, currentLevel)
		queue = newQueue
	}

	return levels
}
