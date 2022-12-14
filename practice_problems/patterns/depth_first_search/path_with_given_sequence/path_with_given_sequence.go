package practice_problems

import . "github.com/ethanjweiner/dsa-go/structures"

// Recursive approach

// func findPath(bt *BT[int], sequence []int) bool {
// 	if len(sequence) == 0 {
// 		return false
// 	}

// 	firstValue := sequence[0]

// 	if bt.Left == nil && bt.Right == nil {
// 		return len(sequence) == 1 && firstValue == bt.Data
// 	}

// 	restOfSequence := sequence[1:]

// 	return firstValue == bt.Data &&
// 		((bt.Left != nil && findPath(bt.Left, restOfSequence)) ||
// 			(bt.Right != nil && findPath(bt.Right, restOfSequence)))
// }

// More procedural

func findPath(bt *BT[int], sequence []int) bool {
	return findPathAcc(bt, sequence, 0)
}

func findPathAcc(bt *BT[int], sequence []int, sequenceIndex int) bool {
	if sequenceIndex >= len(sequence) || bt.Data != sequence[sequenceIndex] {
		return false
	}

	// At this point, we know data equals sequence value

	if bt.Left == nil && bt.Right == nil {
		return sequenceIndex == len(sequence)-1
	}

	return ((bt.Left != nil && findPathAcc(bt.Left, sequence, sequenceIndex+1)) ||
		(bt.Right != nil && findPathAcc(bt.Right, sequence, sequenceIndex+1)))
}
