

class Solution {
    public boolean canPartition(int[] nums) {
        var sum = 0;
        for(int i:nums){
            sum+=i;
        }
        if((sum&1)==1) {
            return false;
        }

        var target = sum/2;

        Boolean[][] cache = new Boolean[nums.length+1][target+1];

        return dfs(nums, nums.length, target, cache);
    }


    public boolean dfs(int[] nums, int n, int target, Boolean[][] cache) {
        if (target == 0 ) {
            return true;
        }

        if (n == 0 || target < 0 ) {
            return false;
        }

        if (cache[n][target] ==null ) {
            cache[n][target] = dfs(nums, n-1,target-nums[n-1],cache) ||dfs(nums, n-1,target,cache);
        }

        return cache[n][target];
    }
}