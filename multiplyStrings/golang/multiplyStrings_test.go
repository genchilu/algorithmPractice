package multiplyStrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func multiplyNum(num []int, multiplicand int) []int {
// func TestDev1(t *testing.T) {
// 	num := []int{9, 9, 9}
// 	multiplicand := 9

// 	expect := []int{8, 9, 9, 1}

// 	actual := multiplyNum(num, multiplicand)

// 	assert.Equal(t, expect, actual)
// }

// //func addNum(a, b []int) []int {
// func TestDev2(t *testing.T) {
// 	num1 := []int{7, 3, 8}
// 	num2 := []int{}

// 	expect := []int{7, 3, 8}

// 	actual := addNum(num1, num2)

// 	assert.Equal(t, expect, actual)

// }

// //func multiply(num1 string, num2 string) string {
// func TestDev3(t *testing.T) {
// 	num1 := "2"
// 	num2 := "3"

// 	expect := "6"

// 	actual := multiply(num1, num2)

// 	assert.Equal(t, expect, actual)
// }

func TestDev4(t *testing.T) {
	num1 := "123"
	num2 := "456"

	expect := "56088"

	actual := multiply(num1, num2)

	assert.Equal(t, expect, actual)
}
