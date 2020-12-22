package addTwoNumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

func TestAddTwoNumbers(t *testing.T) {
	testCases := []struct {
		l1     *ListNode
		l2     *ListNode
		expect *ListNode
	}{
		{
			prepareListNode([]int{2, 4, 3}),
			prepareListNode([]int{5, 6, 4}),
			prepareListNode([]int{7, 0, 8}),
		},
		{
			prepareListNode([]int{0}),
			prepareListNode([]int{0}),
			prepareListNode([]int{0}),
		},
		{
			prepareListNode([]int{9, 9, 9, 9, 9, 9, 9}),
			prepareListNode([]int{9, 9, 9, 9}),
			prepareListNode([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
		{
			prepareListNode([]int{2, 4, 9}),
			prepareListNode([]int{5, 6, 4, 9}),
			prepareListNode([]int{7, 0, 4, 0, 1}),
		},
	}

	for _, testCase := range testCases {
		actual := addTwoNumbers(testCase.l1, testCase.l2)
		expect := testCase.expect
		//showList(actual)
		//showList(expect)
		for expect != nil && actual != nil {
			assert.Equal(t, expect.Val, actual.Val)
			expect = expect.Next
			actual = actual.Next
		}
		// showList(actual)
		// showList(expect)

		var expectNil *ListNode
		assert.Equal(t, expectNil, expect)
		assert.Equal(t, expectNil, actual)
	}
}

func prepareListNode(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	l := &ListNode{nums[0], nil}
	c := l
	for i := 1; i < len(nums); i++ {
		c.Next = &ListNode{nums[i], nil}
		c = c.Next
	}

	return l
}
