import java.util.*;
import java.util.stream.Collector;
import java.util.stream.Collectors;
import java.util.stream.IntStream;

public class KargerMinCut {

    public static int doMinCut(int[][] edges) {
        if (edges == null || edges.length == 0) {
            return 0;
        }

        int vertexCount = countVertices(edges);
        int minCut = edges.length;
        for(int i = 0;i <vertexCount*vertexCount; i++){
            int cutMinCut = kargerMinCut(edges);
            if(cutMinCut < minCut) {
                minCut = cutMinCut;
            }
        }
        return minCut;
    }

    private static int kargerMinCut(int[][] edges) {
        List<List<Integer>> cloneEdges = new ArrayList<>();
        Map<Integer, Integer> vertexMergeIdxMap = new HashMap<>();
        Map<Integer, List<Integer>> mergeIdxVertexMap = new HashMap<>();
        int maxVertex = 0;

        for (int[] edge: edges) {
            Integer v0 = new Integer(edge[0]);
            Integer v1 = new Integer(edge[1]);

            if (vertexMergeIdxMap.get(v0) == null) {
                vertexMergeIdxMap.put(v0, v0);
            }
            if (vertexMergeIdxMap.get(v1) == null) {
                vertexMergeIdxMap.put(v1, v1);
            }

            if (mergeIdxVertexMap.get(v0) == null) {
                List<Integer> mergeIdxGroup = new ArrayList<>();
                mergeIdxGroup.add(v0);
                mergeIdxVertexMap.put(v0, mergeIdxGroup);
            }
            if (mergeIdxVertexMap.get(v1) == null) {
                List<Integer> mergeIdxGroup = new ArrayList<>();
                mergeIdxGroup.add(v1);
                mergeIdxVertexMap.put(v1, mergeIdxGroup);
            }

            if (edge[0] > maxVertex) {
                maxVertex = edge[0];
            }

            if (edge[1] > maxVertex) {
                maxVertex = edge[1];
            }

            cloneEdges.add(
                    Arrays.stream(edge)
                            .boxed()
                            .collect(Collectors.toList())
            );
        }

        int mergeCount = 0;
        while (mergeCount<vertexMergeIdxMap.size()-2) {
            Random r = new Random();
            int pickedIdx = r.nextInt(cloneEdges.size());
            List<Integer> pickledEdge = cloneEdges.get(pickedIdx);

            Integer v0 = pickledEdge.get(0);
            Integer mergeIdx0 = vertexMergeIdxMap.get(v0);

            Integer v1 = pickledEdge.get(1);
            Integer mergeIdx1 = vertexMergeIdxMap.get(v1);

            if(mergeIdx0.intValue() != mergeIdx1.intValue()) {
                mergeCount++;
                Integer curMergeIdx = new Integer(maxVertex + mergeCount);

                List<Integer> mergeIdxGroup0 = mergeIdxVertexMap.get(mergeIdx0);
                mergeIdxGroup0.forEach( v -> vertexMergeIdxMap.put(v, curMergeIdx));

                List<Integer> mergeIdxGroup1 = mergeIdxVertexMap.get(mergeIdx1);
                mergeIdxGroup1.forEach( v -> vertexMergeIdxMap.put(v, curMergeIdx));

                mergeIdxVertexMap.remove(mergeIdx0);
                mergeIdxVertexMap.remove(mergeIdx1);

                mergeIdxGroup0.addAll(mergeIdxGroup1);
                mergeIdxVertexMap.put(curMergeIdx, mergeIdxGroup0);
            }

            cloneEdges.remove(pickedIdx);
        }

        int minCut = 0;
        for (List<Integer> edge: cloneEdges) {
            Integer v0 = edge.get(0);
            Integer mergeIdx0 = vertexMergeIdxMap.get(v0);

            Integer v1 = edge.get(1);
            Integer mergeIdx1 = vertexMergeIdxMap.get(v1);

            if(mergeIdx0.intValue() != mergeIdx1.intValue()) {
                minCut++;
            }
        }
        return minCut;
    }

    private static int countVertices(int[][] edges) {
        Set<Integer> set = new HashSet<>();

        for(int[] edge: edges) {
            set.add(new Integer(edge[0]));
            set.add(new Integer(edge[1]));
        }

        return set.size();
    }
}
