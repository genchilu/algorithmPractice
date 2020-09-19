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

	t := make([]int, 26)
	
	for _, b := range tasks {
		t[b-'A']++
	}

	sort.Ints(t)
	maxT := t[25]
	idle := (maxT-1) * n
	
	for i:=24;i>=0&&idle>0;i-- {
		if(t[i] == t[25]) {
			idle -= (maxT-1)
		} else {
			idle -= t[i]
		}
	}

	if (idle < 0 ) {
		idle = 0
	}

    return len(tasks) + idle
}