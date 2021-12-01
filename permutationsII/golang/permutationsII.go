package permutationsII

func permuteUnique(nums []int) [][]int {

	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	results := [][]int{}
	mypermute(m, []int{}, len(nums), &results)

	return results
}

func mypermute(m map[int]int, p []int, l int, results *[][]int) {
	if (len(p)) == l {
		result := make([]int, l)
		copy(result, p)
		*results = append(*results, result)
	}

	for k, _ := range m {
		if m[k] != 0 {
			m[k]--
			mypermute(m, append(p, k), l, results)
			m[k]++
		}
	}

}
