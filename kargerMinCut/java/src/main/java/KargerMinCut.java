import java.util.*;

public class KargerMinCut {

    public static int doMinCut(int[][] edges) {

        if(edges == null || edges.length == 0) {
            return 0;
        }

        int verticesCount = countVerticesCount(edges);
        int minCut = edges.length;
        for(int i = 0;i<verticesCount*verticesCount;i++) {

            int tmpMinCut = kargerMinCut(edges);
            if (tmpMinCut < minCut) {
                minCut = tmpMinCut;
            }
        }
        return minCut;
    }

    private static int kargerMinCut(int[][] edges) {

        List<List<Integer>> cloneEdges = new ArrayList<>();
        for(int[] edge: edges) {
            List<Integer> e = new ArrayList<>();
            e.add(Integer.valueOf(edge[0]));
            e.add(Integer.valueOf(edge[1]));
            cloneEdges.add(e);
        }

        Map<Integer, Integer> vertexMergeIdxMap = new HashMap<>();
        Map<Integer, List<Integer>> mergeIdxVertexMap = new HashMap<>();

        int maxVertex = initVertexMergeIdxMap(vertexMergeIdxMap, mergeIdxVertexMap, cloneEdges);
        int mergeCount = 0;

        Random rand = new Random();
        while (mergeCount < (vertexMergeIdxMap.size() -2)) {
            int pickIdx = rand.nextInt(cloneEdges.size());
            List<Integer> pickEdges = cloneEdges.get(pickIdx);
            Integer v0 = pickEdges.get(0);
            Integer v1 = pickEdges.get(1);
            Integer mergeIdx0 = vertexMergeIdxMap.get(v0);
            Integer mergeIdx1 = vertexMergeIdxMap.get(v1);

            if(mergeIdx0.intValue() != mergeIdx1.intValue()) {
                mergeCount++;
                Integer curMergeIdx = Integer.valueOf(mergeCount+maxVertex);

                for (Integer v: mergeIdxVertexMap.get(mergeIdx0)) {
                    vertexMergeIdxMap.put(v, curMergeIdx);
                }
                for (Integer v: mergeIdxVertexMap.get(mergeIdx1)) {
                    vertexMergeIdxMap.put(v, curMergeIdx);
                }

                List<Integer> mergeIdxGroup = new ArrayList<>();
                mergeIdxGroup.addAll(mergeIdxVertexMap.get(mergeIdx0));
                mergeIdxGroup.addAll(mergeIdxVertexMap.get(mergeIdx1));
                mergeIdxVertexMap.put(curMergeIdx, mergeIdxGroup);

                mergeIdxVertexMap.remove(mergeIdx0);
                mergeIdxVertexMap.remove(mergeIdx1);
            }

            cloneEdges.remove(pickIdx);
        }

        int minCut = 0;
        for (List<Integer> e: cloneEdges) {
            Integer v0 = e.get(0);
            Integer v1 = e.get(1);

            Integer mergeIdx0 = vertexMergeIdxMap.get(v0);
            Integer mergeIdx1 = vertexMergeIdxMap.get(v1);
            if(mergeIdx0.intValue() != mergeIdx1.intValue()) {
                minCut++;
            }
        }
        return minCut;
    }

    private static int initVertexMergeIdxMap(Map<Integer, Integer> vertexMergeIdxMap, Map<Integer, List<Integer>> mergeIdxVertexMap, List<List<Integer>> edges) {
        int maxVertex = 0;
        for (List<Integer> e: edges) {
            Integer v0 = e.get(0);
            Integer v1 = e.get(1);
            if(vertexMergeIdxMap.get(v0) == null) {
                vertexMergeIdxMap.put(v0, v0);
            }
            if(vertexMergeIdxMap.get(v1) == null) {
                vertexMergeIdxMap.put(v1, v1);
            }

            if(mergeIdxVertexMap.get(v0) == null) {
                mergeIdxVertexMap.put(v0, new ArrayList<>());
                mergeIdxVertexMap.get(v0).add(v0);
            }
            if(mergeIdxVertexMap.get(v1) == null) {
                mergeIdxVertexMap.put(v1, new ArrayList<>());
                mergeIdxVertexMap.get(v1).add(v1);
            }

            if (v0.intValue() > maxVertex) {
                maxVertex = v0.intValue();
            }

            if (v1.intValue() > maxVertex) {
                maxVertex = v1.intValue();
            }
        }

        return maxVertex;

    }

    private static int countVerticesCount(int[][] edges) {
        Set<Integer> verticesSet = new HashSet<>();

        for(int[] e: edges) {
            verticesSet.add(Integer.valueOf(e[0]));
            verticesSet.add(Integer.valueOf(e[1]));
        }

        return verticesSet.size();
    }
}
