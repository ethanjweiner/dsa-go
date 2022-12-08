package practice_problems

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type QuicksortTestSuite struct {
	suite.Suite
	emptySlice []int
	slice1     []int
	slice2     []int
	slice3     []int
}

func (suite *QuicksortTestSuite) SetupTest() {
	suite.slice1 = []int{0, 5, 2, 1, 6, 3}
	suite.slice2 = []int{6, 5}
	suite.slice3 = []int{5, 6}
	suite.emptySlice = []int{}
}

// Tests
func (suite *QuicksortTestSuite) TestPartition() {
	Partition(suite.slice1, 0, len(suite.slice1)-1)
	suite.Equal([]int{0, 1, 2, 3, 6, 5}, suite.slice1)
}

func (suite *QuicksortTestSuite) TestSmallPartition() {
	Partition(suite.slice2, 0, len(suite.slice2)-1)
	suite.Equal([]int{5, 6}, suite.slice2)

	Partition(suite.slice3, 0, len(suite.slice3)-1)
	suite.Equal([]int{5, 6}, suite.slice3)
}

func (suite *QuicksortTestSuite) TestQuicksort() {
	Quicksort(suite.slice1)
	suite.Equal([]int{0, 1, 2, 3, 5, 6}, suite.slice1)
}

func (suite *QuicksortTestSuite) TestFindNthLowest() {
	suite.Equal(FindNthLowest(suite.slice1, 1), 0)
	suite.Equal(FindNthLowest(suite.slice1, 2), 1)
	suite.Equal(FindNthLowest(suite.slice1, 3), 2)
	suite.Equal(FindNthLowest(suite.slice1, 4), 3)
	suite.Equal(FindNthLowest(suite.slice1, 5), 5)
	suite.Equal(FindNthLowest(suite.slice1, 6), 6)
}

func TestQuicksortTestSuite(t *testing.T) {
	suite.Run(t, new(QuicksortTestSuite))
}
