class Solution {
    public int rob(int[] nums) {
        int[] cache = new int[nums.length];
        for(int i=0;i<cache.length;i++){
            cache[i] = -1;
        }

        int plan1 = nums[0] + robDp(nums, 2, nums.length-2,cache);

        for(int i=0;i<cache.length;i++){
            cache[i] = -1;
        }
        int plan2 = robDp(nums, 1,nums.length-1, cache);
        return Math.max(plan1, plan2);
    }

    public int robDp(int[] nums, int start, int end, int[] cache) {
        if(start > end) {
            return 0;
        }

        if(start == end) {
            return nums[start];
        }

        if(cache[start] ==-1) {

            int plan1 = nums[start] + robDp(nums, start + 2, end, cache);
            int plan2 = robDp(nums, start + 1, end, cache);
            cache[start] = Math.max(plan1, plan2);
        }

        return cache[start];
    }
}
