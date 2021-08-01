package maximumProductSubarray

import (
	"math"
)

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}
	max := math.MinInt64

	ps := []int{}
	neg := false

	for _, num := range nums {
		if num == 0 {
			p := cacul(ps, neg)
			if p > max {
				max = p
			}
			if 0 > max {
				max = 0
			}

			ps = []int{}
			neg = false
		} else if num > 0 {
			if len(ps) == 0 {
				ps = append(ps, num)
			} else {
				l := len(ps) - 1
				if ps[l] < 0 {
					ps = append(ps, num)
				} else {
					ps[l] *= num
				}
			}
		} else if num < 0 {
			ps = append(ps, num)
			neg = !neg
		}
		//fmt.Printf("0000 %v\n", ps)
	}

	//fmt.Printf("%v\n", ps)
	p := cacul(ps, neg)

	if p > max {
		return p
	}

	return max
}

func cacul(ps []int, neg bool) int {
	if len(ps) == 0 {
		return 0
	} else if len(ps) == 1 {
		return ps[0]
	} else {
		p := 1
		for _, v := range ps {
			p *= v
		}

		if neg {
			p1 := p / ps[0]
			if ps[0] > 0 {
				p1 /= ps[1]
			}

			p2 := p / ps[len(ps)-1]
			if ps[len(ps)-1] > 0 {
				p2 /= ps[len(ps)-2]
			}

			if p1 > p2 {
				return p1
			} else {
				return p2
			}
		}

		return p
	}
}
