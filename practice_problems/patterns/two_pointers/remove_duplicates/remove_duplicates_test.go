package practice_problems

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type RemoveDuplicatesTestSuite struct {
	suite.Suite
	emptySlice      []int
	oneElementSlice []int
	slice1          []int
	slice2          []int
}

func (suite *RemoveDuplicatesTestSuite) SetupTest() {
	suite.emptySlice = []int{}
	suite.oneElementSlice = []int{2}
	suite.slice1 = []int{2, 3, 3, 3, 6, 9, 9}
	suite.slice2 = []int{2, 2, 2, 11}
}

// Tests
func (suite *RemoveDuplicatesTestSuite) TestGenericCases() {
	suite.slice1 = removeDuplicates(suite.slice1)
	suite.Equal([]int{2, 3, 6, 9}, suite.slice1)

	suite.slice2 = removeDuplicates(suite.slice2)
	suite.Equal([]int{2, 11}, suite.slice2)
}

func (suite *RemoveDuplicatesTestSuite) TestLowLengthSlice() {
	suite.Equal(0, len(removeDuplicates(suite.emptySlice)))
	suite.Equal(1, len(removeDuplicates(suite.oneElementSlice)))
}

func TestRemoveDuplicatesTestSuite(t *testing.T) {
	suite.Run(t, new(RemoveDuplicatesTestSuite))
}
