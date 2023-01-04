package practice_problems

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type BinarySearchTestSuite struct {
	suite.Suite
	emptyArr        []int
	oneElementArray []int
	twoElementArray []int
	oddLengthArray  []int
	evenLengthArray []int
}

func (suite *BinarySearchTestSuite) SetupTest() {
	suite.emptyArr = []int{}
	suite.oneElementArray = []int{3}
	suite.twoElementArray = []int{3, 5}
	suite.oddLengthArray = []int{1, 2, 3, 4, 5}     // end = 4, start = 0, middle = floor(4/2) = 5
	suite.evenLengthArray = []int{1, 2, 3, 4, 5, 6} // end = 5, start = 0, middle = floor(5/2) = 5
}

// Tests
func (suite *BinarySearchTestSuite) TestBinarySearch() {
	suite.Equal(-1, IndexOf(suite.emptyArr, 0))
	suite.Equal(-1, IndexOf(suite.oneElementArray, 0))
	suite.Equal(0, IndexOf(suite.twoElementArray, 3))
	suite.Equal(1, IndexOf(suite.twoElementArray, 5))
	suite.Equal(0, IndexOf(suite.oneElementArray, 3))
	suite.Equal(1, IndexOf(suite.oddLengthArray, 2))
	suite.Equal(2, IndexOf(suite.oddLengthArray, 3))
	suite.Equal(4, IndexOf(suite.oddLengthArray, 5))
	suite.Equal(2, IndexOf(suite.evenLengthArray, 3))
	suite.Equal(3, IndexOf(suite.evenLengthArray, 4))
}

func TestBinarySearchTestSuite(t *testing.T) {
	suite.Run(t, new(BinarySearchTestSuite))
}
