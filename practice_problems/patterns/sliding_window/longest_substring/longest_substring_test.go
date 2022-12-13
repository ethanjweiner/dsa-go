package practice_problems

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type LongestSubstringTestSuite struct {
	suite.Suite
}

func (suite *LongestSubstringTestSuite) SetupTest() {
}

// Tests
func (suite *LongestSubstringTestSuite) TestGenericCases() {
	suite.Equal(4, longestSubstring("araaci", 2))
	suite.Equal(2, longestSubstring("araaci", 1))
	suite.Equal(5, longestSubstring("cbbebi", 3))
	suite.Equal(7, longestSubstring("cbbebbbbi", 2))

}
func (suite *LongestSubstringTestSuite) TestWholeStringMeetsCriteria() {
	suite.Equal(4, longestSubstring("araa", 2))
	suite.Equal(6, longestSubstring("cbbebi", 4))
	suite.Equal(6, longestSubstring("cbbebi", 5))
}

func (suite *LongestSubstringTestSuite) TestKGreaterThanStringLength() {
	suite.Equal(2, longestSubstring("ab", 3))
}

func (suite *LongestSubstringTestSuite) TestEmptyString() {
	suite.Equal(0, longestSubstring("", 3))
}

func TestLongestSubstringTestSuite(t *testing.T) {
	suite.Run(t, new(LongestSubstringTestSuite))
}
