package missingElementInSortedArray

func missingElement(nums []int, k int) int {

	e := nums[0] + k
	for i := 1; i < len(nums); i++ {
		if nums[i] <= e {
			e++
		} else {
			break
		}
	}
	return e
}
