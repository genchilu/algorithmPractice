package longestAbsoluteFilePath

import (
	"strings"
)

func lengthLongestPath(input string) int {
	if len(input) == 0 {
		return 0
	}

	max := 0
	accum := []string{}
	curIdx := 0
	level := 0
	for i := 0; i < len(input); i++ {
		if input[i] == '\n' {
			curPath := input[curIdx:i]
			//fmt.Printf("?? %v, %s\n", accum, curPath)
			if level == len(accum) {
				accum = append(accum, curPath)
			} else {
				if len(accum[level]) <= len(curPath) {
					accum[level] = curPath
				}
			}

			if strings.Contains(curPath, ".") {
				//fmt.Printf("%v\n", accum)
				c := filelength(accum)
				if c > max {
					max = c
				}
			}

			preLevel := level
			curIdx, level, _ = count(input, i)
			//fmt.Printf("222 %d\n", level)

			if level <= preLevel {
				accum = accum[0:level]
			}

			i = curIdx

		} else if i == len(input)-1 {
			curPath := input[curIdx : i+1]
			if level == len(accum) {
				accum = append(accum, curPath)
			} else {
				if len(accum[level]) <= len(curPath) {
					accum[level] = curPath
				}
			}

			if strings.Contains(curPath, ".") {
				//fmt.Printf("%v\n", accum)
				c := filelength(accum)
				if c > max {
					max = c
				}
			}
		}
	}

	return max
}

func count(input string, idx int) (newidx, tc, nc int) {
	for idx < len(input) {
		if input[idx] == '\n' {
			nc++
			idx++
		} else if input[idx] == '\t' {
			tc++
			idx++
		} else {
			break
		}
	}

	return idx, tc, nc
}

func filelength(accum []string) int {
	c := 0
	for j := 0; j < len(accum); j++ {
		c += len(accum[j])
		c += 1
	}
	c--

	return c
}
