package singleElementInASortedArray

func singleNonDuplicate(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	m := len(nums) / 2
	if isOne(nums, m) {
		return nums[m]
	}

	if m&1 == 1 {
		if nums[m] == nums[m-1] {
			return singleNonDuplicate(nums[m+1:])
		} else {
			return singleNonDuplicate(nums[0:m])
		}
	} else {
		if nums[m] == nums[m+1] {
			return singleNonDuplicate(nums[m+2:])
		} else {
			return singleNonDuplicate(nums[0 : m-1])
		}
	}
}

func isOne(nums []int, idx int) bool {
	if idx == 0 {
		return nums[idx+1] != nums[idx]
	}

	if idx == len(nums)-1 {
		return nums[idx-1] != nums[idx]
	}

	return nums[idx+1] != nums[idx] && nums[idx-1] != nums[idx]
}
