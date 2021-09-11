package RangeSumQuery2D

type NumMatrix struct {
	m [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m := make([][]int, len(matrix))
	for i, _ := range matrix {
		accumulation := 0
		m[i] = make([]int, len(matrix[i]))
		for j, _ := range matrix[i] {
			accumulation += matrix[i][j]
			m[i][j] = accumulation
		}
	}

	return NumMatrix{m}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := 0
	for i := row1; i <= row2; i++ {
		sum += this.m[i][col2]
		if col1 != 0 {
			sum -= this.m[i][col1-1]
		}
	}
	return sum
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
