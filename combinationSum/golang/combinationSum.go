package combinationSum

func combinationSum(candidates []int, target int) [][]int {

	result := [][]int{}

	for i, v := range candidates {
		t := target
		pre := []int{}
		for t >= v {
			pre = append(pre, v)
			if v == t {
				result = append(result, pre)
			} else if t > v {
				if i < len(candidates)-1 {
					tmp := combinationSum(candidates[i+1:], t-v)
					for j := range tmp {
						tmp[j] = append(tmp[j], pre...)
					}
					result = append(result, tmp...)
				}
			}

			t -= v
		}
	}

	return result
}
