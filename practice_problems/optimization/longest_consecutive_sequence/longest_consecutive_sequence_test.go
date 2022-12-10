package practice_problems

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type LongestConsecutiveSequenceTestSuite struct {
	suite.Suite
	slice1 []int
	slice2 []int
}

func (suite *LongestConsecutiveSequenceTestSuite) SetupTest() {
	suite.slice1 = []int{10, 5, 12, 3, 55, 30, 4, 11, 2}
	suite.slice2 = []int{19, 13, 15, 12, 18, 14, 17, 11}
}

// Tests
func (suite *LongestConsecutiveSequenceTestSuite) TestLongestConsecutiveSequence() {
	suite.Equal(LongestConsecutiveSequence(suite.slice1), 4)
	suite.Equal(LongestConsecutiveSequence(suite.slice2), 5)
}

func TestLongestConsecutiveSequenceTestSuite(t *testing.T) {
	suite.Run(t, new(LongestConsecutiveSequenceTestSuite))
}
