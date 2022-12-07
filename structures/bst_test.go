package structures

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type BSTTestSuite struct {
	suite.Suite
	leaf1 *BST[int]
	leaf2 *BST[int]
	leaf3 *BST[int]
	bst1  *BST[int]
	bst2  *BST[int]
	bst3  *BST[int]
}

func (suite *BSTTestSuite) SetupTest() {
	suite.leaf1 = &BST[int]{Data: 5}
	suite.leaf2 = &BST[int]{Data: 9}
	suite.leaf3 = &BST[int]{Data: 15}
	suite.bst1 = &BST[int]{Data: 6, Left: suite.leaf1, Right: suite.leaf2}
	suite.bst2 = &BST[int]{Data: 11, Right: suite.leaf3}
	suite.bst3 = &BST[int]{Data: 10, Left: suite.bst1, Right: suite.bst2}
}

// Tests
func (suite *BSTTestSuite) TestInsertOnLeaf() {
	suite.leaf1.Insert(7)
	suite.Equal(7, suite.leaf1.Right.Data)
}

func (suite *BSTTestSuite) TestInsertOnTree() {
	suite.bst3.Insert(7)
	suite.Equal(7, suite.bst3.Left.Right.Left.Data)
}

func (suite *BSTTestSuite) TestInsertOnTreeDuplicate() {
	suite.leaf3.Insert(5)
	suite.Equal(5, suite.bst3.Left.Left.Data)

	// Doesn't duplicate insert data
	suite.Nil(suite.bst3.Left.Left.Left)
	suite.Nil(suite.bst3.Left.Left.Right)
}

func (suite *BSTTestSuite) TestSearchLeaf() {
	suite.Equal(true, suite.leaf1.Search(5))
	suite.Equal(false, suite.leaf1.Search(6))

}

func (suite *BSTTestSuite) TestSearchTree() {
	suite.Equal(true, suite.bst3.Search(5))
	suite.Equal(true, suite.bst3.Search(11))
}

func TestBSTTestSuite(t *testing.T) {
	suite.Run(t, new(BSTTestSuite))
}
