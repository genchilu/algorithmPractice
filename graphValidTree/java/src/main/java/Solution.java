import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

class Solution {
    public boolean validTree(int n, int[][] edges) {
        if(edges.length == 0 ){
            if (n==1) {
                return true;
            }
            return false;
        }
        Map<Integer, List<Integer>> edgesMap = new HashMap<>();
        for(var edge:edges) {
            if(!edgesMap.containsKey(edge[0])) {
                edgesMap.put(edge[0], new ArrayList<>());
            }
            if(!edgesMap.containsKey(edge[1])) {
                edgesMap.put(edge[1], new ArrayList<>());
            }

            edgesMap.get(edge[0]).add(edge[1]);
            edgesMap.get(edge[1]).add(edge[0]);
        }

        boolean[] seen = new boolean[n];
        var res = CheckNoCycle(edgesMap, -1,edges[0][0], seen);

        if(!res.nocycle) {
            return false;
        }

        if(res.nodeCount == n) {
            return true;
        }

        return false;
    }

    class CheckNoCycleRes {
        int nodeCount;
        boolean nocycle;

        public CheckNoCycleRes(int nodeCount, boolean nocycle) {
            this.nodeCount = nodeCount;
            this.nocycle = nocycle;
        }
    }

    public CheckNoCycleRes CheckNoCycle(Map<Integer, List<Integer>> edgesMap,int pre, int cur, boolean[] seen) {
        if(seen[cur]) {
            return new CheckNoCycleRes(-1, false);
        }

        CheckNoCycleRes res = new CheckNoCycleRes(1, true);
        seen[cur] = true;
        for(int next:edgesMap.get(cur)) {
            if(next!=pre) {
                var subRes = CheckNoCycle(edgesMap, cur,next, seen);
                if (!subRes.nocycle) {
                    return new CheckNoCycleRes(-1, false);
                } else {
                    res.nodeCount += subRes.nodeCount;
                }
            }
        }

        return res;

    }
}