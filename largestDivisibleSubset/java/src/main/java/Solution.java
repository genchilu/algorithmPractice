import java.util.*;

class Solution {
    public List<Integer> largestDivisibleSubset(int[] nums) {

        HashMap<Integer, List<Integer>> cache = new HashMap<>();
        Arrays.sort(nums);


        List<Integer> res = new ArrayList<>();
        for(int i =0;i<nums.length;i++){
            List<Integer> set = largestDivisibleSubsetAtIdxK(nums, i, cache);
            if(set.size()>res.size()) {
                res = set;
            }
        }
        return res;
    }

    public List<Integer> largestDivisibleSubsetAtIdxK(int[] nums, int k, HashMap<Integer, List<Integer>> cache) {

        if(cache.containsKey(k)) {
            return cache.get(k);
        }

        List<Integer> max = new ArrayList<>();
        for(int i=0;i<k;i++){
            if(nums[k] % nums[i] == 0) {
                List<Integer> set = largestDivisibleSubsetAtIdxK(nums, i, cache);
                if(set.size() > max.size()) {
                    max = set;
                }
            }
        }

        List<Integer> res = new ArrayList<>();
        res.addAll(max);
        res.add(nums[k]);

        cache.put(k, res);
        return cache.get(k);
    }

}