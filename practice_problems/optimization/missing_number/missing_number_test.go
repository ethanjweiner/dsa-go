package practice_problems

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type MissingNumberTestSuite struct {
	suite.Suite
	slice1 []int
}

func (suite *MissingNumberTestSuite) SetupTest() {
	suite.slice1 = []int{5, 2, 4, 1, 0}
}

// Tests
func (suite *MissingNumberTestSuite) TestMissingNumber() {
	suite.Equal(3, MissingNumber(suite.slice1))
}

func TestMissingNumberTestSuite(t *testing.T) {
	suite.Run(t, new(MissingNumberTestSuite))
}
