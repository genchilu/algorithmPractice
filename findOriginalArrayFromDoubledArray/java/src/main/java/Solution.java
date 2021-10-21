import java.util.*;

class Solution {
    public int[] findOriginalArray(int[] changed) {
        if ((changed.length & 1) == 1) {
            return new int[0];
        }

        Arrays.sort(changed);
        Map<Integer, Integer> map = new HashMap<>();

        List<Integer> result = new ArrayList<>();
        for(int i:changed) {
            if ((i&1) ==1) {
                map.put(i, map.getOrDefault(i, 0)+1);
            } else {
                if (map.getOrDefault(i/2,0) == 0 ){
                    map.put(i, map.getOrDefault(i,0)+1);
                } else {
                    map.put(i/2, map.get(i/2)-1);
                    result.add(i/2);
                }
            }
        }

        for(Integer v: map.values()) {
            if (v!=0) {
                return new int[0];
            }
        }

        int[] origin = result.stream().mapToInt(i->i).toArray();
        return origin;
    }
}