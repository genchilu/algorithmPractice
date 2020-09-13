package randomPickIndex

import(
	"math/rand"
	//"math"
	//"math/bits"
	//"fmt"
)
type Solution struct {
	nums []int
}


func Constructor(nums []int) Solution {


    return Solution{nums}
}


func (this *Solution) Pick(target int) int {

	idxs := []int{}
	for i,v:=range this.nums {
		if v==target {
			idxs = append(idxs, i)
		}
	}

	pick := rand.Intn(len(idxs))
	return idxs[pick]
}