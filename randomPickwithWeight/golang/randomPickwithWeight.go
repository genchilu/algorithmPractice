package randomPickwithWeight

import (
    "math/rand"
)

type Solution struct {
    sumW []int
    sum int
}

func Constructor(w []int) Solution {
    sum:=0
    sumW := make([]int, len(w))
    for i, v := range w {
        sum +=v
        sumW[i] = w[i]
        
    }
    return Solution{sumW,sum}
}


func (this *Solution) PickIndex() int {
    pick := rand.Intn(this.sum)
    for i, v := range this.sumW {
        if v > pick {
            return i
        }
    }

    return (len(this.sumW) - 1)
}
