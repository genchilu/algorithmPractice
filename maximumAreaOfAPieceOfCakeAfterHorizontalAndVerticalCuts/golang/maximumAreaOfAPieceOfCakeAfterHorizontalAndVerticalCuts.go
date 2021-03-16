package maximumAreaOfAPieceOfCakeAfterHorizontalAndVerticalCuts

import (
	"sort"
)

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)
	hsplit := []int{0}
	hsplit = append(hsplit, horizontalCuts...)
	hsplit = append(hsplit, h)

	maxh := 0
	for i := 1; i < len(hsplit); i++ {
		d := hsplit[i] - hsplit[i-1]
		if d > maxh {
			maxh = d
		}
	}

	vsplit := []int{0}
	vsplit = append(vsplit, verticalCuts...)
	vsplit = append(vsplit, w)

	maxv := 0
	for i := 1; i < len(vsplit); i++ {
		d := vsplit[i] - vsplit[i-1]
		if d > maxv {
			maxv = d
		}
	}

	return (maxh * maxv) % 1000000007
}
