package continuousSubarraySum

func checkSubarraySum(nums []int, k int) bool {
	if nums == nil || len(nums) <=1{
		return false
	}

	for i:=0;i<len(nums)-1;i++ {
		sum := nums[i]

		for j:=i+1;j<len(nums);j++ {
			sum += nums[j]
			if k == 0 {
				if sum == 0 {
					return true
				}
			} else if (sum>=k && sum%k==0) {
				return true
			} else if (sum == 0) {
				return true
			}
		}
	}

	return false
}

