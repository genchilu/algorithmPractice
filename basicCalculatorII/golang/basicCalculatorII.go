package basicCalculatorII

import (
	"strconv"
	"strings"
	//"fmt"
)
//var intMap = make(map[string]int)
func calculate(s string) int {
	s=strings.ReplaceAll(s, " ", "")
	//intMap := make(map[string]int)
	return doCalculate(s)
}

func doCalculate(s string) int {
	if strings.Contains(s, "+") || strings.Contains(s, "-") || strings.Contains(s, "*") || strings.Contains(s, "/") {
		idx := strings.LastIndex(s, "+")
		if(idx >= 0) {
			return calculate(s[:idx]) + calculate(s[idx+1:])
		}

		idx = strings.LastIndex(s, "-")
		if(idx >= 0) {
			return calculate(s[:idx]) - calculate(s[idx+1:])
		}

		multipIdx := strings.LastIndex(s, "*")
		divisiIdx := strings.LastIndex(s, "/")

		if(multipIdx >= 0 && divisiIdx >= 0) {
			if (multipIdx > divisiIdx) {
				if calculate(s[multipIdx+1:]) == 0 {
					return 0
				}
				return calculate(s[:multipIdx]) * calculate(s[multipIdx+1:])
			} else {
				if calculate(s[:divisiIdx]) == 0 {
					return 0
				}
				return calculate(s[:divisiIdx]) / calculate(s[divisiIdx+1:])
			}
		} else if multipIdx >= 0 {
			return calculate(s[:multipIdx]) * calculate(s[multipIdx+1:])
		} else {
			return calculate(s[:divisiIdx]) / calculate(s[divisiIdx+1:])
		}
	} else {
		v, _ := strconv.Atoi(s)
		return v
	}
}