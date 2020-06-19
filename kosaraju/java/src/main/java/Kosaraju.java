import java.util.*;
import java.util.stream.Collectors;

public class Kosaraju {
    public Kosaraju(){}

    public Map<Integer, List<Integer>> doKosarajuScc(int[][] inputEdges) {
        Map<Integer, List<Integer>> g = new HashMap<>();

        for (int[] edge: inputEdges) {
            Integer from = new Integer(edge[0]);
            Integer to = new Integer(edge[1]);

            if (g.get(from) == null) {
                g.put(from, new ArrayList<>());
            }

            if (g.get(to) == null) {
                g.put(to, new ArrayList<>());
            }

            g.get(from).add(to);
        }

        Map<Integer, Boolean> isVistited = new HashMap<>();

        Stack<Integer> finishStack = new Stack();
        for(Integer vertex: g.keySet()) {
            dfsRound1(g, vertex, isVistited, finishStack);
        }

        Map<Integer, List<Integer>> revG = reverseG(g);
        isVistited.clear();

        Map<Integer, List<Integer>> componments = new HashMap<>();
        while (!finishStack.isEmpty()) {
            List<Integer> componment = new ArrayList<>();
            dfsRound2(revG, finishStack.pop(), isVistited, componment);

            if(componment.size()>0) {
                Collections.sort(componment);

                Integer lowestVertex = componment.get(0);
                componments.put(lowestVertex, componment);
            }
        }

        return componments;
    }

    private void dfsRound1(Map<Integer, List<Integer>> g, Integer vertex, Map<Integer, Boolean> isVistited, Stack<Integer> finishStack) {
        if(isVistited.get(vertex) == null) {
            isVistited.put(vertex, new Boolean(true));
            List<Integer> toVertices = g.get(vertex);
            for(Integer toVertex: toVertices) {
                dfsRound1(g, toVertex, isVistited, finishStack);
            }
            finishStack.push(vertex);
        }
    }

    private Map<Integer, List<Integer>> reverseG(Map<Integer, List<Integer>> g) {
        Map<Integer, List<Integer>> revG = new HashMap<>();

        for (Integer fromVertex: g.keySet()) {
            List<Integer> toVertices = g.get(fromVertex);

            for(Integer toVertex: toVertices) {
                if (revG.get(toVertex) == null) {
                    revG.put(toVertex, new ArrayList<>());
                }

                if (revG.get(fromVertex) == null) {
                    revG.put(fromVertex, new ArrayList<>());
                }

                revG.get(toVertex).add(fromVertex);
            }
        }

        return revG;
    }

    private void dfsRound2(Map<Integer, List<Integer>> g, Integer vertex, Map<Integer, Boolean> isVistited, List<Integer> componment) {
        if(isVistited.get(vertex) == null) {
            isVistited.put(vertex, new Boolean(true));
            List<Integer> toVertices = g.get(vertex);

            for(Integer toVertex: toVertices) {
                dfsRound2(g, toVertex, isVistited, componment);
            }

            componment.add(vertex);
        }
    }
}
