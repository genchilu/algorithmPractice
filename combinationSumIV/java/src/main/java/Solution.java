import java.util.HashMap;
import java.util.Map;

class Solution {
    Map<Integer, Integer> cache = new HashMap<>();
    public int combinationSum4(int[] nums, int target) {

        //Map<Integer, Integer> cache = new HashMap<>();
        return combinationSum4WithCache(nums, target);
    }

    public int combinationSum4WithCache(int[] nums, int target) {
        if(target==0) {
            return 1;
        }
        if(cache.containsKey(target)) {
            return cache.get(target);
        }

        int count =0;
        for(int num:nums) {
            if(num <= target) {
                count += combinationSum4(nums, target - num);
            }
        }

        cache.put(target, count);
        return cache.get(target);
    }
}