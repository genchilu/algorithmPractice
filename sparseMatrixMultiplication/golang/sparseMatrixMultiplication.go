package sparseMatrixMultiplication

func multiply(A [][]int, B [][]int) [][]int {
	result := make([][]int, len(A))
	for i:=0;i<len(A);i++ {
		result[i] = make([]int, len(B[0]))
		for j:=0;j<len(B[0]);j++ {
			p := 0
			for k:=0;k<len(A[i]);k++ {
				if A[i][k] != 0 && B[k][j] != 0 {
					p += A[i][k]*B[k][j]
				}
			}
			result[i][j] = p
		}
	}
    return result
}