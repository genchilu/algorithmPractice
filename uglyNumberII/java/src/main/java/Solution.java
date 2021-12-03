import java.util.Arrays;
import java.util.HashSet;
import java.util.Set;

class Solution {
    public int nthUglyNumber(int n) {

        int[] result = new int[n];
        result[0]=1;
        int idx5 = 0;
        int idx3 = 0;
        int idx2 = 0;

        for (int i = 1;i<n;i++){
            int candidate2 = result[idx2] * 2;
            int candidate3 = result[idx3] * 3;
            int candidate5 = result[idx5] * 5;
            result[i] = Math.min(candidate2, Math.min(candidate3, candidate5));

            if (result[i] == candidate2) {
                idx2++;
            }
            if (result[i] == candidate3) {
                idx3++;
            }
            if (result[i]==candidate5) {
                idx5++;
            }
        }

        return result[n-1];
    }
}