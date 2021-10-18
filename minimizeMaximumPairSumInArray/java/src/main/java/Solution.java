import java.util.Arrays;

class Solution {
    public int minPairSum(int[] nums) {
        Arrays.sort(nums);
        int max = 0;
        for(int i = 0;i<nums.length/2;i++){
            int t = nums[i]+nums[nums.length-i-1];

            if (t>max) {

                max = t;
            }
        }

        return max;
    }
}