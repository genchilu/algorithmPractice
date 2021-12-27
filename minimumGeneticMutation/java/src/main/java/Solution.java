import java.util.*;

class Solution {
    public int minMutation(String start, String end, String[] bank) {
        if(start.equals(end)) {
            return 0;
        }

        if(bank.length==0) {
            return -1;
        }

        Set<String> seen = new LinkedHashSet<>();
        return dfs(start, end, bank, seen);
    }

    public int dfs(String g, String target, String[] bank, Set<String> seen) {
        seen.add(g);

        int min = Integer.MAX_VALUE;
        for(String nextG: bank) {
            if(isMutation(g, nextG)) {
                if(nextG.equals(target)) {
                    min = 1;
                } else {

                    if (!seen.contains(nextG)) {
                        int count = dfs(nextG, target, bank, seen) + 1;
                        if (count != 0) {
                            if (count < min) {
                                min = count;
                            }
                        }
                    }
                }
            }
        }

        seen.remove(g);

        if(min==Integer.MAX_VALUE) {
            return -1;
        }

        return min;
    }
    public boolean isMutation(String a, String b) {

        int count = 0;
        for(int i=0;i<a.length();i++){
            if(a.charAt(i) != b.charAt(i)) {
                count++;
            }
        }

        if(count==1) {
            return true;
        }

        return false;
    }
}