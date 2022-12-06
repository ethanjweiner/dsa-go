package structures

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type StackTestSuite struct {
	suite.Suite
	stack1 *Stack
}

func (suite *StackTestSuite) SetupTest() {
	suite.stack1 = &Stack{}
	suite.stack1.Push(1)
	suite.stack1.Push(2)
	suite.stack1.Push(3)
}

// Tests
func (suite *StackTestSuite) TestPop() {
	suite.Equal(3, suite.stack1.Pop())
	suite.Equal(2, suite.stack1.Pop())
	suite.Equal(1, suite.stack1.Pop())
	suite.Nil(suite.stack1.Pop())
}

func TestStackTestSuite(t *testing.T) {
	suite.Run(t, new(StackTestSuite))
}
