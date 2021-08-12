package countOfSmallerNumbersAfterSelf

func countSmaller(nums []int) []int {

	count := make([]int, len(nums))
	indices := make([]int, len(nums))

	for i := range indices {
		indices[i] = i
	}

	l, r := 0, len(nums)-1

	merge(nums, l, r, &indices, &count)
	return count
}

func merge(nums []int, l, r int, indices, count *[]int) {
	if l >= r {
		return
	}

	m := (l + r) / 2
	merge(nums, l, m, indices, count)
	merge(nums, m+1, r, indices, count)

	i, j := l, m+1
	tmp := []int{}
	for i <= m && j <= r {
		if nums[(*indices)[i]] <= nums[(*indices)[j]] {
			tmp = append(tmp, (*indices)[i])
			(*count)[(*indices)[i]] += j - m - 1
			i++
		} else {
			tmp = append(tmp, (*indices)[j])
			j++
		}
	}

	for i <= m {
		(*count)[(*indices)[i]] += j - m - 1
		tmp = append(tmp, (*indices)[i])
		i++
	}

	for j <= m+1 {
		tmp = append(tmp, (*indices)[j])
		j++
	}

	for k := range tmp {
		(*indices)[l+k] = tmp[k]
	}

}
