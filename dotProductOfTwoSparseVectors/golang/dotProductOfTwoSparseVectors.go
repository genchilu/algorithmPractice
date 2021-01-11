package dotProductOfTwoSparseVectors

type SparseVector struct {
	m map[int]int
}

func Constructor(nums []int) SparseVector {
	m := make(map[int]int)

	for i, v := range nums {
		if v != 0 {
			m[i] = v
		}
	}
	return SparseVector{m}
}

// Return the dotProduct of two sparse vectors
func (this *SparseVector) dotProduct(vec SparseVector) int {
	sum := 0
	for k, v := range this.m {
		if vv, ok := vec.m[k]; ok {
			sum += v * vv
		}
	}
	return sum
}
