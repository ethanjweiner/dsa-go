package structures

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type QueueTestSuite struct {
	suite.Suite
	queue1 *Queue
}

func (suite *QueueTestSuite) SetupTest() {
	suite.queue1 = &Queue{}
	suite.queue1.Enqueue(1)
	suite.queue1.Enqueue(2)
	suite.queue1.Enqueue(3)
}

// Tests
func (suite *QueueTestSuite) TestDequeue() {
	suite.Equal(1, suite.queue1.Dequeue())
	suite.Equal(2, suite.queue1.Dequeue())
	suite.Equal(3, suite.queue1.Dequeue())
	suite.Nil(suite.queue1.Dequeue())
}

func TestQueueTestSuite(t *testing.T) {
	suite.Run(t, new(QueueTestSuite))
}
