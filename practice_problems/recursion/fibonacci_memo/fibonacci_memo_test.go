package practice_problems

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type FibMemoTestSuite struct {
	suite.Suite
}

// Tests
func (suite *FibMemoTestSuite) TestFibMemo() {
	suite.Equal(0, Fib(0))
	suite.Equal(1, Fib(1))
	suite.Equal(1, Fib(2))
	suite.Equal(2, Fib(3))
	suite.Equal(3, Fib(4))
	suite.Equal(5, Fib(5))
	suite.Equal(55, Fib(10))
}

func TestFibMemoTestSuite(t *testing.T) {
	suite.Run(t, new(FibMemoTestSuite))
}
