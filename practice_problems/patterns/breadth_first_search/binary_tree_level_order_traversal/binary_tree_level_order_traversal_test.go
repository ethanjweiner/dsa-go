package practice_problems

import (
	"testing"

	. "github.com/ethanjweiner/dsa-go/structures"
	"github.com/stretchr/testify/suite"
)

type BinaryTreeLevelOrderTraversalTestSuite struct {
	suite.Suite
	bt1 *BT[int]
	bt2 *BT[int]
	bt3 *BT[int]
	bt4 *BT[int]
	bt5 *BT[int]
	bt6 *BT[int]
}

/*
BT1:

     1
   /   \
  2     3
 / \   / \
4  5  6   7

BT4:
     12
	 /    \
	7     1
 /     / \
9    10  5
*/

func (suite *BinaryTreeLevelOrderTraversalTestSuite) SetupTest() {
	suite.bt2 = &BT[int]{Data: 2, Left: &BT[int]{Data: 4}, Right: &BT[int]{Data: 5}}
	suite.bt3 = &BT[int]{Data: 3, Left: &BT[int]{Data: 6}, Right: &BT[int]{Data: 7}}
	suite.bt1 = &BT[int]{Data: 1, Left: suite.bt2, Right: suite.bt3} // primary
	suite.bt5 = &BT[int]{Data: 7, Left: &BT[int]{Data: 9}}
	suite.bt6 = &BT[int]{Data: 1, Left: &BT[int]{Data: 10}, Right: &BT[int]{Data: 5}}
	suite.bt4 = &BT[int]{Data: 12, Left: suite.bt5, Right: suite.bt6} // primary
}

// Tests
func (suite *BinaryTreeLevelOrderTraversalTestSuite) TestTraverse() {
	suite.ElementsMatch([][]int{{5}}, traverse(&BT[int]{Data: 5}))
	suite.ElementsMatch([][]int{{1}, {2, 3}, {4, 5, 6, 7}}, traverse(suite.bt1))
	suite.ElementsMatch([][]int{{12}, {7, 1}, {9, 10, 5}}, traverse(suite.bt4))
}

func TestBinaryTreeLevelOrderTraversalTestSuite(t *testing.T) {
	suite.Run(t, new(BinaryTreeLevelOrderTraversalTestSuite))
}
