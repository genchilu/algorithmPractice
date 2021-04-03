package containerWithMostWater

func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	max := 0

	for l < r {
		a := 0
		if height[l] < height[r] {
			a = height[l] * (r - l)
			l++
		} else {
			a = height[r] * (r - l)
			r--
		}
		if a > max {
			max = a
		}
	}
	return max
}
