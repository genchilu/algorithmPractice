import java.util.*;

public class Kosaraju {
    private class Graph {
        private Map<Integer, List<Integer>> edges;

        public Graph(){
            edges = new HashMap<>();
        }

        public void addEdge(int from, int to) {
            Integer fromInteger = new Integer(from);
            Integer toInteger = new Integer(to);

            List<Integer> toVertices = edges.get(fromInteger);
            if (toVertices == null) {
                toVertices = new ArrayList<>();
            }
            toVertices.add(toInteger);
            edges.put(fromInteger, toVertices);

            if(edges.get(toInteger) == null) {
                edges.put(toInteger, new ArrayList<>());
            }
        }

        public List<Integer> getNeighboList(Integer vertexI) {
            return edges.get(vertexI);
        }

        public Graph getReverseGraph() {
            Graph g = new Graph();
            for (Map.Entry<Integer, List<Integer>> entry: edges.entrySet()) {
                int from = entry.getKey().intValue();
                for (Integer toI: entry.getValue()) {
                    int to = toI.intValue();
                    g.addEdge(to, from);
                }
            }

            return g;
        }

        public Set<Integer> getAllVertices() {
            return this.edges.keySet();
        }
    }

    public Kosaraju(){}

    public Map<Integer, List<Integer>> doKosarajuScc(int[][] inputEdges) {

        Graph g = new Graph();
        for (int[] edge: inputEdges) {
            g.addEdge(edge[0], edge[1]);
        }

        Stack finishStack = new Stack();
        Map<Integer, Boolean> isVisitedMap = new HashMap<>();
        for(Integer vertex: g.getAllVertices()) {
            dfs1(g, vertex, isVisitedMap, finishStack);
        }

        Graph gt = g.getReverseGraph();
        Map<Integer, List<Integer>> components = new HashMap<>();

        Map<Integer, Boolean> isVisited2Map = new HashMap<>();
        while(!finishStack.isEmpty()) {
            Integer curVertex  = (Integer) finishStack.pop();
            List<Integer> componment = new ArrayList<>();

            dfs2(gt, curVertex, isVisited2Map, componment);

            if(componment.size() > 0) {
                Collections.sort(componment);

                Integer lowestVertex = componment.get(0);
                components.put(lowestVertex, componment);
            }
        }

        return components;
    }

    private void dfs1(Graph g, Integer currentVertex, Map<Integer, Boolean> visitedMap, Stack finishStack) {
        Boolean isVisited = visitedMap.get(currentVertex);
        if(isVisited == null || !isVisited.booleanValue()) {
            visitedMap.put(currentVertex, Boolean.TRUE);

            List<Integer> allNeighborList = g.getNeighboList(currentVertex);
            for (Integer neighbor: allNeighborList) {
                dfs1(g, neighbor, visitedMap, finishStack);
            }

            finishStack.push(currentVertex);
        }
    }

    private void dfs2(Graph g, Integer currentVertex, Map<Integer, Boolean> visitedMap, List<Integer> componment) {
        Boolean isVisited = visitedMap.get(currentVertex);
        if(isVisited == null || !isVisited.booleanValue()) {
            visitedMap.put(currentVertex, Boolean.TRUE);
            componment.add(currentVertex);
            List<Integer> allNeighborList = g.getNeighboList(currentVertex);
            for (Integer neighbor: allNeighborList) {
                dfs2(g, neighbor, visitedMap, componment);
            }
        }
    }

}
