package kClosestPointsToOrigin

import "sort"

func KClosest(points [][]int, K int) [][]int {
	if len(points) == 0 {
		return [][]int{}
	}

	dmap := make(map[int][][]int)
	ds := make([]int, len(points))
	for i, p:=range points {
		ds[i] = p[0] * p[0] + p[1] * p[1]
		if _,ok:=dmap[ds[i]];ok{
			dmap[ds[i]] = append(dmap[ds[i]], p)
		} else {
			dmap[ds[i]] = [][]int{p}
		}
	}
	sort.Ints(ds)

	result := make([][]int, 0, K)
	notbreak := true
	for notbreak {
		for _, v:= range ds {
			for _,vv:=range dmap[v] {
				result = append(result, vv)
				if len(result) == K {
					notbreak = false
					break
				}
			}
			if !notbreak {
				break
			}
		}
	}

    return result
}