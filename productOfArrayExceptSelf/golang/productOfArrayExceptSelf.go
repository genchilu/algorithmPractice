package productOfArrayExceptSelf

import "math"
//import "fmt"

func productExceptSelf(nums []int) []int {
	if nums == nil || len(nums) == 0{
		return []int{}
	}

	countMap := make(map[int]int)
	intMap := make(map[int]int)

	for _, v := range nums {
		if _, ok := countMap[v];!ok {
			countMap[v] = 0
			intMap[v] = 1
		}
		countMap[v]++
	}

	//p *= (int)(math.Pow((float64)(i), (float64)(countMap[c])))
	if _, ok := countMap[0];ok {
		if countMap[0] >= 2 {
			for i, _ := range nums {
				nums[i] = 0
			}

			return nums
		}

		p := 0
		for i, c := range countMap {
			if i != 0 {
				if p == 0 {
					p = 1
				}
				p *= (int) (math.Pow((float64)(i), (float64)(c)))
			} 
		}
		for i, _ := range countMap {
			if (i != 0) {
				intMap[i] = 0
			}
		}
		intMap[0] = p
	} else {
		for i, c := range countMap {
			for j, v := range intMap {
				if i == j {
					tmp := 1
					if i!=1 {
						tmp = (int)(math.Pow((float64)(i), (float64)(c-1)))
					}
					intMap[j] = v * tmp
				} else {
					tmp := 1
					if i!=1 {
						tmp = (int)(math.Pow((float64)(i), (float64)(c)))
					}
					intMap[j] = v * tmp
				}
			}
		}
	}

	for i, v := range nums {
		nums[i] = intMap[v]
	}


    return nums
}