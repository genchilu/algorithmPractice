import java.util.*;

class Solution {
    public String removeDuplicateLetters(String s) {
        int[] occurs = new int[26];
        List<Byte> stack = new ArrayList<>();

        boolean[] seen = new boolean[26];

        byte[] bs = s.getBytes();
        for(int i =0;i<bs.length;i++){
            occurs[bs[i]- 'a']++;
        }

        for(int i =0;i<bs.length;i++) {
            if(seen[bs[i]-'a']) {
                occurs[bs[i]-'a']--;
                continue;
            }

            if(stack.size() == 0) {
                stack.add(bs[i]);
                seen[bs[i]-'a'] = true;
                occurs[bs[i]-'a']--;
                continue;
            }

            if(stack.get(stack.size()-1) < bs[i]) {
                stack.add(bs[i]);
                seen[bs[i]-'a'] = true;
                occurs[bs[i]-'a']--;
                continue;
            }

            while(stack.size()>0 && stack.get(stack.size()-1) > bs[i] && occurs[stack.get(stack.size()-1)-'a'] > 0) {
                byte b = stack.get(stack.size()-1);
                stack.remove(stack.size()-1);
                seen[b-'a'] = false;
            }

            stack.add(bs[i]);
            seen[bs[i]-'a'] = true;
            occurs[bs[i]-'a']--;
        }

        byte[] res = new byte[stack.size()];
        for(int i = 0;i<stack.size();i++){
            res[i] = stack.get(i);
        }
        return new String(res);
    }
}