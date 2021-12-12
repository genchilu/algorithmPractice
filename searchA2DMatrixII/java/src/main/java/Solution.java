import java.util.Arrays;
import java.util.Comparator;

class Solution {

    public boolean searchMatrix(int[][] matrix, int target) {

        if (matrix.length==0 || matrix[0].length==0) {
            return false;
        }

        for(int i=0;i<matrix.length;i++){
            var res = Arrays.binarySearch(matrix[i], target);
            if (res >= 0) {
                return true;
            }
        }
        return false;
    }

}