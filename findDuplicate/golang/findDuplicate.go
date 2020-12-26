package findDuplicate

func findDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[slow]

	for fast != slow {
		fast = nums[nums[fast]]
		slow = nums[slow]
	}

	slow = 0

	for slow != fast {
		fast = nums[fast]
		slow = nums[slow]
	}

	return slow
}
