package practice_problems

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type SubArraysWithProductTestSuite struct {
	suite.Suite
	slice1 []int
	slice2 []int
}

func (suite *SubArraysWithProductTestSuite) SetupTest() {
	suite.slice1 = []int{2, 5, 3, 10}
	suite.slice2 = []int{8, 2, 6, 5}
}

// Tests
func (suite *SubArraysWithProductTestSuite) TestSubArrayWithProduct() {
	result := subarrayProductLessThanK(suite.slice1, 30)
	suite.ElementsMatch([][]int{{2}, {5}, {2, 5}, {3}, {5, 3}, {10}}, result)
}

func TestSubArraysWithProductTestSuite(t *testing.T) {
	suite.Run(t, new(SubArraysWithProductTestSuite))
}
