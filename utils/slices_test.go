package utils

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type SlicesTestSuite struct {
	suite.Suite
	slice1 []int
}

func (suite *SlicesTestSuite) SetupTest() {
	suite.slice1 = []int{1, 2, 3, 4, 5}
}

// Tests
func (suite *SlicesTestSuite) TestDelete() {
	newSlice, _ := Delete(suite.slice1, 2)
	suite.Equal([]int{1, 2, 4, 5}, newSlice)
}

func (suite *SlicesTestSuite) TestDeleteOutOfBounds() {
	_, err := Delete(suite.slice1, 5)
	suite.Error(err)
}
func (suite *SlicesTestSuite) TestContiguousSubarraysFromStart() {
	suite.ElementsMatch([][]int{{2}, {2, 3}, {2, 3, 4}}, ContiguousSubarraysFromStart(suite.slice1, 1, 4))
}

func (suite *SlicesTestSuite) TestContiguousSubarraysFromEnd() {
	subarrays := ContiguousSubarraysFromEnd(suite.slice1, 1, 4)
	suite.ElementsMatch([][]int{{4}, {3, 4}, {2, 3, 4}}, subarrays)
}

func TestSlicesTestSuite(t *testing.T) {
	suite.Run(t, new(SlicesTestSuite))
}
