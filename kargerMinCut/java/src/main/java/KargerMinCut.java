import java.util.*;

public class KargerMinCut {

    public static int doMinCut(int[][] edges) {
        if (edges == null || edges.length == 0) {
            return 0;
        }

        Set vertexSet = getAllVertex(edges);
        int vertexNum = vertexSet.size();
        int minCut = edges.length;
        for (int i = 0;i<vertexNum*vertexNum;i++) {
            int tmpMinCut = kargerMinCut(edges, vertexSet);
            if(tmpMinCut < minCut) {
                minCut = tmpMinCut;
            }
        }
        return minCut;
    }

    private static Set<Integer> getAllVertex(int[][] edges) {
        Set<Integer> vertexSet = new HashSet<>();
        for (int[] edge: edges) {
            vertexSet.add(new Integer(edge[0]));
            vertexSet.add(new Integer(edge[1]));
        }
        return vertexSet;
    }

    private static int kargerMinCut(int[][] edges, Set vertexSet) {
        Map<Integer, Integer> vertexMergeIdxMap = initVertexMergeMap(vertexSet);
        Map<Integer, List<Integer>> revVertexMergeIdxMap = initRevVertexMergeMap(vertexSet);

        List<Integer[]> edgesList = new ArrayList<>(edges.length);
        for (int[] edge : edges) {
            edgesList.add(new Integer[]{new Integer(edge[0]), new Integer(edge[1])});
        }

        int mergeCount = 0;

        while(mergeCount < (vertexSet.size()-2)) {
            Integer curIdx = new Integer(edges.length + mergeCount);

            int pickIdx = (int) (Math.random()*edgesList.size());
            Integer[] pickedEdge = edgesList.get(pickIdx);

            Integer v1 = pickedEdge[0];
            Integer curV1Idx = vertexMergeIdxMap.get(v1);
            Integer v2 = pickedEdge[1];
            Integer curV2Idx = vertexMergeIdxMap.get(v2);

            if(curV1Idx.intValue() != curV2Idx.intValue()) {

                for (Integer v: revVertexMergeIdxMap.get(curV1Idx)) {
                    vertexMergeIdxMap.put(v, curIdx);
                }

                for (Integer v: revVertexMergeIdxMap.get(curV2Idx)) {
                    vertexMergeIdxMap.put(v, curIdx);
                }


                List<Integer> mergeIdxGroup = new ArrayList<>();
                mergeIdxGroup.addAll(revVertexMergeIdxMap.get(curV1Idx));
                revVertexMergeIdxMap.remove(curV1Idx);
                mergeIdxGroup.addAll(revVertexMergeIdxMap.get(curV2Idx));
                revVertexMergeIdxMap.remove(curV2Idx);

                revVertexMergeIdxMap.put(curIdx, mergeIdxGroup);

                mergeCount++;
            }

            edgesList.remove(pickIdx);
        }

        int minCut = 0;
        for(Integer[] edge: edgesList) {
            if(vertexMergeIdxMap.get(edge[0]).intValue() != vertexMergeIdxMap.get(edge[1]).intValue()) {
                minCut++;
            }
        }

        return minCut;
    }

    private static Map<Integer, Integer> initVertexMergeMap(Set vertexSet) {
        Map<Integer, Integer> vertexMergeMap = new HashMap<>();

        Iterator<Integer> iter = vertexSet.iterator();
        while(iter.hasNext()) {
            Integer vertex = iter.next();
            vertexMergeMap.put(vertex, vertex);
        }

        return vertexMergeMap;
    }

    private static Map<Integer, List<Integer>> initRevVertexMergeMap(Set vertexSet) {
        Map<Integer, List<Integer>> revVertexMergeMap = new HashMap<>();

        Iterator<Integer> iter = vertexSet.iterator();
        while(iter.hasNext()) {
            Integer vertex = iter.next();
            ArrayList<Integer> mergeVertexGroup = new ArrayList<>();
            mergeVertexGroup.add(vertex);
            revVertexMergeMap.put(vertex, mergeVertexGroup);
        }

        return revVertexMergeMap;
    }
}
