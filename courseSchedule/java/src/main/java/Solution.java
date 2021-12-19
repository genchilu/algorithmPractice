import java.util.HashMap;
import java.util.Map;

class Solution {
    public boolean canFinish(int numCourses, int[][] prerequisites) {
        boolean[][] requires = new boolean[numCourses][numCourses];
        Map<Integer, Boolean> cache = new HashMap<>();
        for(var require: prerequisites) {
            requires[require[0]][require[1]] = true;
        }

        for(int i =0;i<numCourses;i++) {
            for(int j=0;j<numCourses;j++){
                if(requires[i][j]) {
                    boolean[] seen = new boolean[numCourses];
                    seen[i] = true;
                    if(!checkNoCycle(seen,j,requires, cache)) {
                        return false;
                    }
                }
            }
        }
        return true;
    }

    public boolean checkNoCycle(boolean[] seen, int cur, boolean[][] requires, Map<Integer, Boolean> cache) {
        if(cache.containsKey(cur)) {
            return cache.get(cur);
        }
        if(seen[cur]) {
            cache.put(cur, false);
            return cache.get(cur);
        }

        seen[cur] = true;
        for(int i =0;i<requires[cur].length;i++) {
            if(requires[cur][i]) {

                if (!checkNoCycle(seen, i, requires, cache)) {
                    cache.put(cur, false);
                    return cache.get(cur);
                }
            }
        }
        seen[cur] = false;

        cache.put(cur, true);
        return cache.get(cur);
    }
}