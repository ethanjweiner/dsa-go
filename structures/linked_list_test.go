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
	Node4        *Node
	Node5        *Node
	Node6        *Node

	Reversed *LinkedList
}

// Initialize fields for setup
func (suite *LinkedListTestSuite) SetupTest() {
	suite.EmptyList = &LinkedList{}
	suite.Node1 = &Node{Data: 1}
	suite.Node2 = &Node{Data: 2}
	suite.Node3 = &Node{Data: 3}
	suite.NonEmptyList = &LinkedList{} // 3 -> 2 - 1
	suite.NonEmptyList.Prepend(suite.Node1)
	suite.NonEmptyList.Prepend(suite.Node2)
	suite.NonEmptyList.Prepend(suite.Node3)

	suite.Node4 = &Node{Data: 1}
	suite.Node5 = &Node{Data: 2}
	suite.Node6 = &Node{Data: 3}
	suite.Reversed = &LinkedList{} // 1 -> 2 -> 3
	suite.Reversed.Prepend(suite.Node6)
	suite.Reversed.Prepend(suite.Node5)
	suite.Reversed.Prepend(suite.Node4)
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
	suite.EmptyList.Prepend(suite.Node2)
	suite.Equal(suite.EmptyList.Head, suite.Node2)
	suite.Equal(suite.EmptyList.Head.Next, suite.Node1)
	suite.EmptyList.Prepend(suite.Node3)
	suite.Equal(suite.EmptyList.Head, suite.Node3)
	suite.Equal(suite.EmptyList.Head.Next, suite.Node2)
	suite.Equal(suite.EmptyList.Head.Next.Next, suite.Node1)

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

func (suite *LinkedListTestSuite) TestReverse() {
	suite.NonEmptyList.Reverse()
	suite.Equal(suite.Reversed, suite.NonEmptyList)
}

func (suite *LinkedListTestSuite) TestReverseEmpty() {
	suite.EmptyList.Reverse()
	suite.Equal(suite.EmptyList, suite.EmptyList)
}

// Run the suite itself w/ normal go test function
func TestLinkedListTestSuite(t *testing.T) {
	suite.Run(t, new(LinkedListTestSuite))
}
