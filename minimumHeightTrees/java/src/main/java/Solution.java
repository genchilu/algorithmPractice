import java.util.*;

class Solution {
    public List<Integer> findMinHeightTrees(int n, int[][] edges) {
        Map<Integer, Set<Integer>> neighborMap = new HashMap<>();
        for(int i=0;i<n;i++){
            neighborMap.put(i, new HashSet<>());
        }

        for(int[] edge: edges) {
            neighborMap.get(edge[0]).add(edge[1]);
            neighborMap.get(edge[1]).add(edge[0]);
        }


        while(neighborMap.size()>2) {
            List<Integer> removeVertices = new ArrayList<>();
            for(int vertex:neighborMap.keySet()) {
                if(neighborMap.get(vertex).size() == 1) {
                    removeVertices.add(vertex);
                }
            }

            for(int vertex1: removeVertices) {
                int vertex2 = neighborMap.get(vertex1).iterator().next();
                neighborMap.remove(vertex1);
                neighborMap.get(vertex2).remove(vertex1);
            }

        }

        List<Integer> res = new ArrayList<>();
        for(int vertex:neighborMap.keySet()) {
            res.add(vertex);
        }

        return res;
    }


}