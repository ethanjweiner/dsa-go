package structures

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

// Set up suite & suite fields
type LinkedListTestSuite struct {
	suite.Suite
	EmptyList    *LinkedList
	NonEmptyList *LinkedList
	Node1        *Node
	Node2        *Node
	Node3        *Node
}

// Initialize fields for setup
func (suite *LinkedListTestSuite) SetupTest() {
	suite.EmptyList = &LinkedList{}
	suite.Node1 = &Node{Data: 1}
	suite.Node2 = &Node{Data: 2}
	suite.Node3 = &Node{Data: 3}
	suite.NonEmptyList = &LinkedList{Head: suite.Node1}
	suite.NonEmptyList.Prepend(suite.Node2)
	suite.NonEmptyList.Prepend(suite.Node3)
}

func (suite *LinkedListTestSuite) TestSearch() {
	suite.True(suite.NonEmptyList.Search(1))
	suite.True(suite.NonEmptyList.Search(2))
	suite.True(suite.NonEmptyList.Search(3))
	suite.False(suite.NonEmptyList.Search(0))
}

func (suite *LinkedListTestSuite) TestPrepend() {
	suite.EmptyList.Prepend(suite.Node1)
	suite.Equal(suite.EmptyList.Head, suite.Node1)

	suite.Equal(suite.NonEmptyList.Head, suite.Node3)
	suite.Equal(suite.NonEmptyList.Head.Next, suite.Node2)
}

func (suite *LinkedListTestSuite) TestDeleteWithValueHead() {
	initialLength := suite.NonEmptyList.Length
	suite.NonEmptyList.DeleteWithValue(3)
	suite.Equal(suite.NonEmptyList.Length, initialLength-1)
	suite.Equal(suite.NonEmptyList.Head, suite.Node2)
}

func (suite *LinkedListTestSuite) TestDeleteWithValueMiddle() {
	initialLength := suite.NonEmptyList.Length
	suite.NonEmptyList.DeleteWithValue(2)
	suite.Equal(suite.NonEmptyList.Length, initialLength-1)
	suite.Equal(suite.NonEmptyList.Head, suite.Node3)
}

func (suite *LinkedListTestSuite) TestDeleteWithValueTail() {
	initialLength := suite.NonEmptyList.Length
	suite.NonEmptyList.DeleteWithValue(1)
	suite.Equal(suite.NonEmptyList.Length, initialLength-1)
	suite.Equal(suite.NonEmptyList.Head, suite.Node3)
	suite.Nil(suite.NonEmptyList.Head.Next.Next)
}

func (suite *LinkedListTestSuite) TestDeleteWithValueFromEmpty() {
	suite.EmptyList.DeleteWithValue(1)
	suite.Equal(suite.EmptyList.Length, 0)
}

func (suite *LinkedListTestSuite) TestDeleteNonExistentValue() {
	initialLength := suite.NonEmptyList.Length
	suite.NonEmptyList.DeleteWithValue(5)
	suite.Equal(suite.NonEmptyList.Length, initialLength)
}

// Run the suite itself w/ normal go test function
func TestLinkedListTestSuite(t *testing.T) {
	suite.Run(t, new(LinkedListTestSuite))
}
