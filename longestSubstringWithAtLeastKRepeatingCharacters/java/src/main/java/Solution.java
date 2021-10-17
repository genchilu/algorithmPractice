import java.util.HashMap;
import java.util.HashSet;
import java.util.Map;
import java.util.Set;

class Solution {
    public int longestSubstring(String s, int k) {
        Set<Character> uniqc = new HashSet<>();
        for(int i =0;i<s.length();i++){
            uniqc.add(s.charAt(i));
        }

        int result = 0;
        for(int i=1;i<=uniqc.size();i++) {
            int l = 0;
            int r = 0;
            int curuniqc = 0;
            int atleastKc = 0;
            int[] bcount = new int[26];

            while(r<s.length()) {
                if (curuniqc <=i) {
                    if (bcount[s.charAt(r)-'a']==0) {
                        curuniqc++;
                    }

                    bcount[s.charAt(r)-'a']++;

                    if(bcount[s.charAt(r)-'a'] == k) {
                        atleastKc++;
                    }

                    r++;
                } else {
                    if(bcount[s.charAt(l)-'a'] == k) {
                        atleastKc--;
                    }

                    bcount[s.charAt(l)-'a']--;

                    if(bcount[s.charAt(l)-'a'] == 0) {
                        curuniqc--;
                    }

                    l++;
                }

                if (curuniqc==i && curuniqc==atleastKc) {
                    result = Math.max(result, r-l);
                }


            }
        }
        return result;
    }
}