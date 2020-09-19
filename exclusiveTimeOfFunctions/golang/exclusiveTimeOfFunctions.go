package exclusiveTimeOfFunctions

import (
	"strings"
	"strconv"
	//"fmt"
)
func exclusiveTime(n int, logs []string) []int {
    if n == 1 {
		_, t, _ := parseLog(logs[len(logs)-1])
		return []int{t+1}
	}

	result := make([]int, n)
	stack := []int{}

	used := 0
	for _, l := range logs {
		id, t, act := parseLog(l)
		if(act == "end") {
			preT := stack[len(stack)-1]
			ft := t - preT + 1 -used
			result[id] += ft
			used += ft
			stack = stack[0:len(stack)-1]
		} else {
			_, t, _ := parseLog(l)
			t = t-used
			stack = append(stack, t)
		}
	}

	return result
}

func parseLog(s string) (id, t int, act string) {
	tmp := strings.Split(s, ":")
	id, _ = strconv.Atoi(tmp[0])
	act = tmp[1]
	t, _ = strconv.Atoi(tmp[2])


	return id,t,act
}