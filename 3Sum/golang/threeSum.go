package threeSum

import "sort"

func threeSum(nums []int) [][]int {
	result := [][]int{}
	sort.Ints(nums)
	for l, _ := range nums {
		if nums[l] > 0 {
			break
		}
		if l == 0 || nums[l] != nums[l-1] {
			result = append(result, twoSum(l, l+1, len(nums)-1, nums)...)
		}
	}
	return result
}

func twoSum(l, m, r int, nums []int) [][]int {
	result := [][]int{}
	if l < m && m < r {
		sum := nums[l] + nums[m] + nums[r]
		if sum < 0 {
			result = append(result, twoSum(l, m+1, r, nums)...)
		} else if sum > 0 {
			result = append(result, twoSum(l, m, r-1, nums)...)
		} else {
			result = append(result, []int{nums[l], nums[m], nums[r]})

			newm := m
			for newm < r {
				if nums[newm] == nums[m] {
					newm++
				} else {
					break
				}
			}
			result = append(result, twoSum(l, newm, r-1, nums)...)
		}
	}
	return result
}
