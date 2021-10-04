package rangeAddition

func getModifiedArray(length int, updates [][]int) []int {

	r := make([]int, length)
	for _, update := range updates {
		for i := update[0]; i <= update[1]; i++ {
			r[i] += update[2]
		}
	}
	return r
}
