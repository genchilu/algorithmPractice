package subsets

func subsets(nums []int) [][]int {

	subsets := [][]int{{}}

	for _, num := range nums {
		for _, subset := range subsets {
			newSubset := append([]int{}, subset...)
			newSubset = append(newSubset, num)
			subsets = append(subsets, newSubset)
		}
	}
	return subsets
}
