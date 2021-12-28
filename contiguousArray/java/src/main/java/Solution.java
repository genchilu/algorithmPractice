import java.util.HashMap;
import java.util.Map;

class Solution {
    public int findMaxLength(int[] nums) {
        Map<Integer, Integer> map = new HashMap<>();
        map.put(0,-1);
        int count =0;
        int max = 0;
        for(int i =0;i<nums.length;i++){
            if(nums[i] == 0) {
                count--;
            } else {
                count++;
            }

            if(map.containsKey(count)) {
                int l = i-map.get(count);
                if(l>max){
                    max = l;
                }
            } else {
                map.put(count, i);
            }
        }


        return max;
    }
}