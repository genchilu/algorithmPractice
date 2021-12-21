import java.util.HashMap;
import java.util.Map;

class Solution {
    public int threeSumSmaller(int[] nums, int target) {
        Map<Integer, Map<Integer, Map<Integer, Integer>>> cache = new HashMap<>();
        return sum(nums, 0, 3, target, cache);
    }

    public int sum(int[] nums, int idx, int count, int target, Map<Integer, Map<Integer, Map<Integer, Integer>>> cache) {

        if(cache.containsKey(idx) &&
                cache.get(idx).containsKey(count) &&
                cache.get(idx).get(count).containsKey(target)) {
            return cache.get(idx).get(count).get(target);
        }

        int res = 0;

        if (count == 1) {
            for(int i = idx;i<nums.length;i++) {
                if (nums[i]<target) {
                    res++;
                }
            }
        } else {

            for (int i = idx; i < nums.length; i++) {
                res += sum(nums, i + 1, count - 1, target - nums[i], cache);
            }
        }

        if (!cache.containsKey(idx)) {
            cache.put(idx, new HashMap<>());
        }
        if(!cache.get(idx).containsKey(count)) {
            cache.get(idx).put(count, new HashMap<>());
        }

        cache.get(idx).get(count).put(target, res);
        return cache.get(idx).get(count).get(target);
    }
}