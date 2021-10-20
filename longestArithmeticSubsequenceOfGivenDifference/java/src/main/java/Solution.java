import java.util.HashMap;
import java.util.Map;

class Solution {
    public int longestSubsequence(int[] arr, int difference) {

        Map<Integer, Integer> map = new HashMap<>();

        for(int n:arr) {
            map.put(n, map.getOrDefault(n-difference,0)+1);
        }

        int max = 0;
        for(int n: map.values()) {
            if (n > max) {
                max = n;
            }
        }

        return max;
    }

}