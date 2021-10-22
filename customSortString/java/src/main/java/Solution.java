import java.util.HashMap;
import java.util.Map;

class Solution {
    public String customSortString(String order, String s) {
        Map<Character, Integer> map = new HashMap<>();
        for (int i =0;i<s.length();i++){
            map.put(s.charAt(i), map.getOrDefault(s.charAt(i), 0)+1);
        }

        char[] cs = new char[s.length()];
        int idx=0;
        for(int i=0;i<order.length();i++){
            for (int j = 0;j<map.getOrDefault(order.charAt(i),0);j++){
                cs[idx]=order.charAt(i);
                idx++;
            }

            map.remove(order.charAt(i));
        }

        for (Character c: map.keySet()) {
            for (int j = 0;j<map.getOrDefault(c,0);j++){
                cs[idx]=c;
                idx++;
            }
        }

        return new String(cs);
    }
}