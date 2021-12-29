import java.util.HashMap;
import java.util.HashSet;
import java.util.Map;
import java.util.Set;

class Solution {
    public int findPairs(int[] nums, int k) {
        Map<Integer, Integer> map = new HashMap<>();
        for(int num:nums) {
            int count = map.getOrDefault(num,0);
            map.put(num, count+1);
        }

        int count = 0;

        if(k==0){
            for(Integer num:map.keySet()) {
                if(map.containsKey(num) && map.get(num) > 1) {
                    count++;
                }
            }

            return count;
        }

        for(Integer num:map.keySet()) {

            if(map.containsKey(num+k)) {
                count++;
            }
            if(map.containsKey(num-k)){
                count++;
            }
        }

        return count/2;
    }
}