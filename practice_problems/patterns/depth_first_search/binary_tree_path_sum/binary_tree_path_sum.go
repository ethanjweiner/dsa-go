package practice_problems

import . "github.com/ethanjweiner/dsa-go/structures"

func hasSumPath(bt *BT[int], sum int) bool {
	if bt == nil {
		return false
	}

	if bt.Left == nil && bt.Right == nil {
		return bt.Data == sum
	}

	nextSum := sum - bt.Data

	return hasSumPath(bt.Left, nextSum) || hasSumPath(bt.Right, nextSum)
}
