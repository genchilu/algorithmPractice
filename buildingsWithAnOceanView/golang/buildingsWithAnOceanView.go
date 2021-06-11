package buildingsWithAnOceanView

func findBuildings(heights []int) []int {

	max := -1
	result := []int{}
	for l := len(heights) - 1; l >= 0; l-- {
		if heights[l] > max {
			result = append(result, l)
			max = heights[l]
		}
	}

	i, j := 0, len(result)-1
	for i < j {
		result[i], result[j] = result[j], result[i]
		i++
		j--
	}

	return result
}
