import java.util.*;
import java.util.stream.Collectors;

public class KargerMinCut {

    public static int doMinCut(int[][] edges) {
        if (edges == null || edges.length == 0){
            return 0;
        }

        int vertexCount =caculateVertexCount(edges);
        int minCut = edges.length;
        for(int i = 0;i< vertexCount*vertexCount*vertexCount;i++) {
            int tmpMinCut = doKargerMinCut(edges);
            if(tmpMinCut < minCut) {
                minCut = tmpMinCut;
            }
        }
        return minCut;
    }


    private static int doKargerMinCut(int[][] edges) {

        List<List<Integer>> cloneEdges = Arrays.asList(edges).stream().map(edge ->
            Arrays.asList(Integer.valueOf(edge[0]), Integer.valueOf(edge[1]))
        ).collect(Collectors.toList());


        Map<Integer, Integer> vertexMergeIdxMap = new HashMap<>();
        Map<Integer, List<Integer>> mergeIdxVertexMap = new HashMap<>();
        int maxVertex = 0;

        for (List<Integer> edge: cloneEdges) {
            Integer v0 = Integer.valueOf(edge.get(0));
            Integer v1 = Integer.valueOf(edge.get(1));

            vertexMergeIdxMap.put(v0, v0);
            vertexMergeIdxMap.put(v1, v1);

            if(mergeIdxVertexMap.get(v0) == null) {
                mergeIdxVertexMap.put(v0, new LinkedList<>());
                mergeIdxVertexMap.get(v0).add(v0);
            }
            if(mergeIdxVertexMap.get(v1) == null) {
                mergeIdxVertexMap.put(v1, new LinkedList<>());
                mergeIdxVertexMap.get(v1).add(v1);
            }

            if(v0.intValue() > maxVertex) {
                maxVertex = v0.intValue();
            }
            if(v1.intValue() > maxVertex) {
                maxVertex = v1.intValue();
            }
        }

        Random r = new Random();
        r.setSeed(System.currentTimeMillis());
        int mergeCount = 0;
        while (mergeCount < vertexMergeIdxMap.size()-2) {
            mergeCount++;
            Integer newMergeIdx = Integer.valueOf(mergeCount + maxVertex);

            int pickIdx = r.nextInt(cloneEdges.size());
            List<Integer> pickEdge = cloneEdges.get(pickIdx);
            Integer v0 = Integer.valueOf(pickEdge.get(0));
            Integer mergeIdx0 = vertexMergeIdxMap.get(v0);
            Integer v1 = Integer.valueOf(pickEdge.get(1));
            Integer mergeIdx1 = vertexMergeIdxMap.get(v1);

            if (mergeIdx0.intValue() != mergeIdx1.intValue()) {
                for(Integer v: mergeIdxVertexMap.get(mergeIdx0)) {
                    vertexMergeIdxMap.put(v, newMergeIdx);
                }
                for(Integer v: mergeIdxVertexMap.get(mergeIdx1)) {
                    vertexMergeIdxMap.put(v, newMergeIdx);
                }

                mergeIdxVertexMap.get(mergeIdx0).addAll(mergeIdxVertexMap.get(mergeIdx1));
                mergeIdxVertexMap.put(newMergeIdx, mergeIdxVertexMap.get(mergeIdx0));

                mergeIdxVertexMap.remove(mergeIdx0);
                mergeIdxVertexMap.remove(mergeIdx1);


            }

            cloneEdges.remove(pickEdge);
        }

        int minCut = 0;

        for(List<Integer> edge: cloneEdges) {
            Integer v0 = Integer.valueOf(edge.get(0));
            Integer mergeIdx0 = vertexMergeIdxMap.get(v0);
            Integer v1 = Integer.valueOf(edge.get(1));
            Integer mergeIdx1 = vertexMergeIdxMap.get(v1);
            if (mergeIdx0.intValue() != mergeIdx1.intValue()) {
                minCut++;
            }
        }

        return minCut;
    }

    private static int caculateVertexCount(int[][] edges) {

        Map<Integer, Boolean> map = new HashMap<>();
        for (int[] edge: edges) {
            Integer v0 = Integer.valueOf(edge[0]);
            Integer v1 = Integer.valueOf(edge[1]);

            map.put(v0, Boolean.TRUE);
            map.put(v1, Boolean.TRUE);
        }

        return map.size();
    }
}
