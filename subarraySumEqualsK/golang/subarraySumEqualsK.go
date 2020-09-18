package subarraySumEqualsK


func subarraySum(nums []int, k int) int {
	if nums == nil {
		return 0
	}

	sumMap := make(map[int]int)
	sum := 0
	count :=0
	for _, v:=range nums {
		sum += v
		if sum == k {
			count++
		}

		if c,ok := sumMap[sum-k];ok{
			count +=c
		}

		if c, ok := sumMap[sum];!ok {
			sumMap[sum] = 1;
		} else {
			sumMap[sum] = 1 + c
		}
	}
    return count
}