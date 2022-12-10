package practice_problems

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type MergeTwoSortedListsTestSuite struct {
	suite.Suite
	l1  *ListNode
	l2  *ListNode
	l12 *ListNode
}

func (suite *MergeTwoSortedListsTestSuite) SetupTest() {
	suite.l1 = &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	suite.l2 = &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	suite.l12 = &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, nil}}}}}}
}

// Tests
func (suite *MergeTwoSortedListsTestSuite) TestMergeTwoSortedLists() {
	suite.Equal(suite.l12, mergeTwoLists(suite.l1, suite.l2))
}

func TestMergeTwoSortedListsTestSuite(t *testing.T) {
	suite.Run(t, new(MergeTwoSortedListsTestSuite))
}
