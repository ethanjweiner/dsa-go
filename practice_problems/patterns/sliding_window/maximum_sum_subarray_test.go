package practice_problems

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type MaximumSumSubarrayTestSuite struct {
	suite.Suite
	slice1 []int
	slice2 []int
}

func (suite *MaximumSumSubarrayTestSuite) SetupTest() {
	suite.slice1 = []int{2, 1, 5, 1, 3, 2}
	suite.slice2 = []int{2, 3, 4, 1, 5}
}

// Tests
func (suite *MaximumSumSubarrayTestSuite) TestMaximumSumSubarray() {
	suite.Equal(9, MaximumSumSubarray(suite.slice1, 3))
	suite.Equal(7, MaximumSumSubarray(suite.slice2, 2))
}

func TestMaximumSumSubarrayTestSuite(t *testing.T) {
	suite.Run(t, new(MaximumSumSubarrayTestSuite))
}
