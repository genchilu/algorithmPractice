package targetSum

func findTargetSumWays(nums []int, target int) int {

	cache := make(map[int]map[int]int)
	return dfs(nums, 0, 0, target, cache)
}

func dfs(nums []int, idx, sum, target int, cache map[int]map[int]int) int {
	if sum == target && idx == len(nums) {
		return 1
	}

	if idx == len(nums) {
		return 0
	}

	if _, ok := cache[idx]; ok {
		if _, ok2 := cache[idx][sum]; ok2 {
			return cache[idx][sum]
		}
	}

	count := 0
	count += dfs(nums, idx+1, sum+nums[idx], target, cache)
	count += dfs(nums, idx+1, sum-nums[idx], target, cache)

	if _, ok := cache[idx]; !ok {
		cache[idx] = make(map[int]int)
	}

	cache[idx][sum] = count
	return cache[idx][sum]
}
