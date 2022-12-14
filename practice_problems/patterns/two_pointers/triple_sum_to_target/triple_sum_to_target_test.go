package practice_problems

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TripleSumToTargetTestSuite struct {
	suite.Suite
	slice1 []int
	slice2 []int
	slice3 []int
	slice4 []int
}

func (suite *TripleSumToTargetTestSuite) SetupTest() {
	suite.slice1 = []int{-2, 0, 1, 2}
	suite.slice2 = []int{-3, -1, 1, 2}
	suite.slice3 = []int{1, 0, 1, 1}
	suite.slice4 = []int{-3, -1, 1, 2, 4}
}

// Tests
func (suite *TripleSumToTargetTestSuite) TestTripleSumToTarget() {
	suite.Equal(1, searchTriplet(suite.slice1, 2))
	suite.Equal(0, searchTriplet(suite.slice2, 1))
	suite.Equal(3, searchTriplet(suite.slice3, 100))
	suite.Equal(3, searchTriplet(suite.slice4, 3))
	suite.Equal(4, searchTriplet(suite.slice4, 4))
	suite.Equal(5, searchTriplet(suite.slice4, 6))
	suite.Equal(7, searchTriplet(suite.slice4, 7))
	suite.Equal(-3, searchTriplet(suite.slice4, -4))
}

func (suite *TripleSumToTargetTestSuite) TestPairSumToTarget() {
	suite.Equal(0, searchPair(suite.slice1, 0, 0))
	suite.Equal(1, searchPair(suite.slice1, 0, 1))
	suite.Equal(2, searchPair(suite.slice3, 100, 0))
	suite.Equal(-4, searchPair(suite.slice4, -5, 0))
	suite.Equal(-2, searchPair(suite.slice2, -2, 0))
	suite.Equal(0, searchPair(suite.slice2, 0, 0))
}

func (suite *TripleSumToTargetTestSuite) TestClosestNumberToTarget() {
	suite.Equal(2, closestNumberToTarget(suite.slice1, 3))
	suite.Equal(-2, closestNumberToTarget(suite.slice1, -3))
	suite.Equal(-3, closestNumberToTarget(suite.slice4, -3))
	suite.Equal(-1, closestNumberToTarget(suite.slice2, 0))
}

func TestTripleSumToTargetTestSuite(t *testing.T) {
	suite.Run(t, new(TripleSumToTargetTestSuite))
}
