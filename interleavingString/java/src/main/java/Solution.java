import java.util.HashMap;
import java.util.Map;

class Solution {
    public boolean isInterleave(String s1, String s2, String s3) {
        if((s1.length()+s2.length()!=s3.length())) {
            return false;
        }
        
        Map<String, Boolean> cache = new HashMap<>();
        return check(s1,s2,s3, cache);
    }

    public boolean check(String s1, String s2, String s3, Map<String, Boolean> cache) {
        if(s1.length()==0) {
            return s2.equals(s3);
        }

        if(s2.length()==0) {
            return s1.equals(s3);
        }

        if(cache.containsKey(s1+","+s2)) {
            return cache.get(s1+","+s2);
        }

        if(s1.charAt(0) == s3.charAt(0)) {
            if(check(s1.substring(1), s2, s3.substring(1), cache)) {
                cache.put(s1+","+s2, true);
                return true;
            }
        }

        if(s2.charAt(0) == s3.charAt(0)) {
            if(check(s1, s2.substring(1), s3.substring(1), cache)) {
                cache.put(s1+","+s2, true);
                return true;
            }
        }

        cache.put(s1+","+s2, false);
        return false;
    }
}
