package practice_problems

import (
	"testing"

	. "github.com/ethanjweiner/dsa-go/structures"
	"github.com/stretchr/testify/suite"
)

type MinimumDepthOfBinaryTreeTestSuite struct {
	suite.Suite
	bt          *BT[int]
	btLeft      *BT[int]
	btRight     *BT[int]
	btRightLeft *BT[int]
}

func (suite *MinimumDepthOfBinaryTreeTestSuite) SetupTest() {
	suite.btLeft = &BT[int]{Data: 7, Left: &BT[int]{Data: 9}}
	suite.btRightLeft = &BT[int]{Data: 10, Right: &BT[int]{Data: 11}}
	suite.btRight = &BT[int]{Data: 1, Left: suite.btRightLeft, Right: &BT[int]{Data: 5}}
	suite.bt = &BT[int]{Data: 12, Left: suite.btLeft, Right: suite.btRight}
}

/*
bt:
      12
     /  \
    7    1
   /     / \
  9     10  5
         \
          11
*/

// Tests
func (suite *MinimumDepthOfBinaryTreeTestSuite) TestFindDepth() {
	suite.Equal(3, findDepth(suite.bt))
}

func TestMinimumDepthOfBinaryTreeTestSuite(t *testing.T) {
	suite.Run(t, new(MinimumDepthOfBinaryTreeTestSuite))
}
