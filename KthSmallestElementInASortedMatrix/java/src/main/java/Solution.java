class Solution {
    public int kthSmallest(int[][] matrix, int k) {

        return binary(matrix, 0, matrix.length*matrix.length-1, k);
    }

    public int binary(int[][] matrix, int l, int r, int k) {
        int pm = l/matrix.length;
        int pn = l%matrix.length;

        int idx = l+1;

        for (int i=l+1;i<=r && idx<=r;i++){
            var m = i/matrix.length;
            var n = i % matrix.length;
            if (matrix[m][n] < matrix[pm][pn]) {
                swap(matrix, idx, i);
                idx++;
            }
        }

        idx--;
        swap(matrix, l, idx);
        if (idx+1 == k) {
            var m = idx/matrix.length;
            var n = idx % matrix.length;
            return matrix[m][n];
        } else if (idx < k) {
            return binary(matrix, idx+1,r,k);
        } else {
            return binary(matrix, l,idx-1,k);
        }
    }

    public void swap(int[][] matrix, int a, int b) {
        int am = a/matrix.length;
        int an = a%matrix.length;

        int bm = b/matrix.length;
        int bn = b%matrix.length;

        int tmp = matrix[am][an];
        matrix[am][an] = matrix[bm][bn];
        matrix[bm][bn] = tmp;
    }
}