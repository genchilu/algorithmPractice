package largestNumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func func largestNumber(nums []int) string {
func Test1(t *testing.T) {
	testCases := []struct {
		a      int
		b      int
		expect bool
	}{
		{
			100,
			9,
			false,
		},
		{
			990,
			9,
			false,
		},
		{
			90,
			8,
			true,
		},
		{
			3,
			34,
			false,
		},
	}

	for _, testCase := range testCases {
		actual := compare(testCase.a, testCase.b)
		assert.Equal(t, testCase.expect, actual)
	}
}

func Test2(t *testing.T) {
	testCases := []struct {
		nums   []int
		expect string
	}{
		{
			[]int{100, 9},
			"9100",
		},
		{
			[]int{990, 9},
			"9990",
		},
		{
			[]int{90, 8},
			"908",
		},
		{
			[]int{10, 2},
			"210",
		},
		{
			[]int{10},
			"10",
		},
		{
			[]int{1},
			"1",
		},
	}

	for _, testCase := range testCases {
		actual := largestNumber(testCase.nums)
		assert.Equal(t, testCase.expect, actual)
	}
}

func Test3(t *testing.T) {
	testCases := []struct {
		nums   []int
		expect string
	}{
		{
			[]int{3, 30, 34, 5, 9},
			"9534330",
		},
	}

	for _, testCase := range testCases {
		actual := largestNumber(testCase.nums)
		assert.Equal(t, testCase.expect, actual)
	}
}

func Test4(t *testing.T) {
	testCases := []struct {
		nums   []int
		expect string
	}{
		{
			[]int{34323, 3432},
			"343234323",
		},
		{
			[]int{0, 0},
			"0",
		},
	}

	for _, testCase := range testCases {
		actual := largestNumber(testCase.nums)
		assert.Equal(t, testCase.expect, actual)
	}
}
