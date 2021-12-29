import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

class Solution {
    int[][] cache;
    public int maximumScore(int[] nums, int[] multipliers) {
        cache = new int[multipliers.length][multipliers.length];
        for(int i =0;i<cache.length;i++){
            Arrays.fill(cache[i], Integer.MIN_VALUE);
        }
        return calculateMax(nums, 0, multipliers, 0);
    }

    public int calculateMax(int[] nums, int li, int[] multipliers, int j) {
        //nums.length-1-ri + li = j
        //ri = nums.length -1 -j + li
        if(j==multipliers.length) {
            return 0;
        }

        int ri = nums.length -1 -j + li;

        if(cache[li][j] != Integer.MIN_VALUE) {
            return cache[li][j];
        }

        int sum = Math.max(
                nums[li] * multipliers[j] + calculateMax(nums, li+1,multipliers,j+1),
                nums[ri] * multipliers[j] + calculateMax(nums, li,multipliers,j+1)
        );

        cache[li][j] = sum;

        return cache[li][j];

    }
}