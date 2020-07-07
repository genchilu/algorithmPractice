import java.awt.*;
import java.util.*;
import java.util.List;

public class Kosaraju {
    public Kosaraju(){}

    public Map<Integer, List<Integer>> doKosarajuScc(int[][] inputEdges) {
        Map<Integer, List<Integer>> g = new HashMap<>();
        for (int[] edge: inputEdges) {
            Integer from = Integer.valueOf(edge[0]);
            Integer to = Integer.valueOf(edge[1]);
            if(g.get(from) == null) {
                g.put(from, new ArrayList<>());
            }
            if(g.get(to) == null) {
                g.put(to, new ArrayList<>());
            }
            g.get(from).add(to);
        }

        Map<Integer, Boolean> isVistited = new HashMap<>();
        Stack<Integer> finishStack = new Stack<>();
        for (Integer v: g.keySet()) {
            dfsRound1(g, v, isVistited, finishStack);
        }

        Map<Integer, List<Integer>> revG = reverseG(g);
        Map<Integer, Boolean> isVistited2 = new HashMap<>();
        Map<Integer, List<Integer>> componments = new HashMap<>();

        while(!finishStack.isEmpty()) {
            List<Integer> componment = new ArrayList<>();
            dfsRound2(revG, finishStack.pop(), isVistited2, componment);

            if(componment.size() >0) {
                Collections.sort(componment);
                componments.put(componment.get(0), componment);
            }
        }


        return componments;
    }


    private static void dfsRound1(Map<Integer, List<Integer>> g, Integer cutV, Map<Integer, Boolean> isVistited, Stack<Integer> finishStack) {
        if (isVistited.get(cutV) == null) {
            isVistited.put(cutV, Boolean.TRUE);

            for (Integer toVs: g.get(cutV)) {
                dfsRound1(g, toVs, isVistited, finishStack);
            }

            finishStack.push(cutV);
        }
    }

    private static Map<Integer, List<Integer>> reverseG(Map<Integer, List<Integer>> g) {
        Map<Integer, List<Integer>> revG = new HashMap<>();

        for(Integer from: g.keySet()) {
            if (revG.get(from) == null) {
                revG.put(from, new ArrayList<>());
            }
            for(Integer to: g.get(from)) {
                if (revG.get(to) == null) {
                    revG.put(to, new ArrayList<>());
                }
                revG.get(to).add(from);
            }
        }

        return revG;
    }

    private static void dfsRound2(Map<Integer, List<Integer>> g, Integer curV, Map<Integer, Boolean> isVistited, List<Integer> componment) {
        if(isVistited.get(curV) == null) {
            isVistited.put(curV, Boolean.TRUE);
            for(Integer v: g.get(curV)) {
                dfsRound2(g, v, isVistited, componment);
            }

            componment.add(curV);
        }
    }
}
