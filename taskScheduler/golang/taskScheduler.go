package taskScheduler

import (
	//"fmt"
	"sort"
)

func leastInterval(tasks []byte, n int) int {
	if len(tasks) == 0 {
		return 0
	}

	if len(tasks) == 1 {
		return 1
	}

	tmp := make(map[byte]int)
	for _, t := range tasks {
		if v, ok := tmp[t];ok {
			tmp[t] = v + 1
		} else {
			tmp[t] = 1
		}
	}

	taskCount := 0
	ts := make([]int, len(tmp))
	c:=0
	for _,v:=range tmp {
		ts[c] = v
		taskCount+=v
		c++
	}

	sort.Sort(sort.Reverse(sort.IntSlice(ts)))

	result := []int{}

	for taskCount>0 {
		for i:=0;i<=n;i++ {
			if i<len(ts) && ts[i] > 0 {
				taskCount--
				ts[i]--
			}

			result = append(result, 0)
			if taskCount == 0 {
				break
			}
		}
		sort.Sort(sort.Reverse(sort.IntSlice(ts)))
	}

    return len(result)
}