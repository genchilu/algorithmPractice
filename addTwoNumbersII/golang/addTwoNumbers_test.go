package addTwoNumbers

import (
	"fmt"
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
			prepareListNode([]int{7, 2, 4, 3}),
			prepareListNode([]int{5, 6, 4}),
			prepareListNode([]int{7, 8, 0, 7}),
		},
	}

	for _, testCase := range testCases {
		actual := addTwoNumbers(testCase.l1, testCase.l2)
		expect := testCase.expect

		showList(actual)
		showList(expect)
		for expect != nil && actual != nil {
			assert.Equal(t, expect.Val, actual.Val)
			expect = expect.Next
			actual = actual.Next
		}

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

func showList(l *ListNode) {
	for l != nil {
		fmt.Printf("%d ", l.Val)
		l = l.Next
	}
	fmt.Printf("\n")
}
