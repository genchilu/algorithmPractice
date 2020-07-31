import java.awt.*;
import java.util.*;
import java.util.List;

public class Kosaraju {
    public Kosaraju(){}

    public Map<Integer, List<Integer>> doKosarajuScc(int[][] inputEdges) {
        if (inputEdges == null || inputEdges.length == 0) {
            return new HashMap<>();
        }

        Map<Integer, List<Integer>> g = new HashMap<>();
        for(int[] edge: inputEdges) {
            Integer from = Integer.valueOf(edge[0]);
            Integer to = Integer.valueOf(edge[1]);

            if (g.get(from) == null) {
                g.put(from, new ArrayList<>());
            }

            if (g.get(to) == null) {
                g.put(to, new ArrayList<>());
            }

            g.get(from).add(to);
        }

        Map<Integer, Boolean> isVistited = new HashMap<>();
        Stack<Integer> finishStack = new Stack<>();
        for(Integer v: g.keySet()) {
            dfs1(g, v, isVistited, finishStack);
        }

        Map<Integer, List<Integer>> revG = reverse(g);

        Map<Integer, List<Integer>> componments = new HashMap<>();
        Map<Integer, Boolean> isVistited2 = new HashMap<>();
        while(!finishStack.isEmpty()) {
            Integer v = finishStack.pop();
            List<Integer> componment = new ArrayList<>();

            dfs2(revG, v, isVistited2, componment);

            if(componment.size()>0) {
                Collections.sort(componment);
                componments.put(componment.get(0), componment);
            }
        }

        return componments;
    }

    private void dfs1(Map<Integer, List<Integer>> g, Integer curVertex, Map<Integer, Boolean> isVistited, Stack finishStack) {
        if (isVistited.get(curVertex) == null) {
            isVistited.put(curVertex, Boolean.TRUE);

            for(Integer v: g.get(curVertex)) {
                dfs1(g, v, isVistited, finishStack);
            }

            finishStack.push(curVertex);
        }
    }

    private Map<Integer, List<Integer>> reverse(Map<Integer, List<Integer>> g) {
        Map<Integer, List<Integer>> revG = new HashMap<>();

        for(Integer from: g.keySet()) {
            for(Integer to: g.get(from)) {
                if(revG.get(to) == null) {
                    revG.put(to, new ArrayList<>());
                }
                revG.get(to).add(from);
            }

            if(revG.get(from) == null) {
                revG.put(from, new ArrayList<>());
            }
        }

        return revG;
    }

    private void dfs2(Map<Integer, List<Integer>> g, Integer curVertex, Map<Integer, Boolean> isVistited, List<Integer> componment) {
        if(isVistited.get(curVertex) == null){
            isVistited.put(curVertex, Boolean.TRUE);

            for(Integer v: g.get(curVertex)) {
                dfs2(g, v, isVistited, componment);
            }

            componment.add(curVertex);
        }

    }
}
