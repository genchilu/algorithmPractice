class Solution {
    public int minSubArrayLen(int target, int[] nums) {
        int l =0;
        int r =0;
        int min = nums.length+1;
        int sum = 0;
        while(sum>= target || r<nums.length) {
            if (sum<target) {
                sum+=nums[r];
                r++;
            } else {
                var tmpMin = r-l;
                if (tmpMin<min) {
                    min = tmpMin;
                }
                sum-=nums[l];
                l++;
            }
        }

        if(min==nums.length+1) {
            return 0;
        }
        return min;
    }
}