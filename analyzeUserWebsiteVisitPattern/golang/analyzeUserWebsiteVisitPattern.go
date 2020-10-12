package analyzeUserWebsiteVisitPattern

import (
	"fmt"
	"sort"
	"strings"
)

func mostVisitedPattern(username []string, timestamp []int, website []string) []string {
	um := make(map[string][]int)
	tu := make(map[int]map[string]string)

	for i, _ := range timestamp {
		user := username[i]
		time := timestamp[i]
		web := website[i]
		if _, ok := um[user]; !ok {
			um[user] = []int{}
		}
		um[user] = append(um[user], time)

		if _, ok := tu[time]; !ok {
			tu[time] = make(map[string]string)
		}

		tu[time][user] = web
	}

	ps := make(map[string]int)

	for u, ts := range um {
		sort.Ints(ts)
		paths := []string{}

		for _, t := range ts {
			paths = append(paths, tu[t][u])

			fmt.Printf("222 %s, %s\n", u, tu[t][u])
		}

		ups := make(map[string]bool)
		for i := 0; i < len(paths)-2; i++ {
			for j := i + 1; j < len(paths)-1; j++ {
				for k := j + 1; k < len(paths); k++ {
					pattern := paths[i] + "|" + paths[j] + "|" + paths[k]
					ups[pattern] = true
				}
			}
		}

		for p, _ := range ups {
			if _, ok := ps[p]; !ok {
				ps[p] = 0
			}
			ps[p]++
		}
	}

	var maxp []string
	maxc := 0
	for p, c := range ps {
		tmp := strings.Split(p, "|")
		if c > maxc {
			maxp = tmp
			maxc = c
		} else if c == maxc {
			for i, _ := range tmp {
				compare := strings.Compare(tmp[i], maxp[i])
				if compare == -1 {
					maxp = tmp
					break
				} else if compare == 1 {
					break
				}
			}
		}
	}

	return maxp
}
