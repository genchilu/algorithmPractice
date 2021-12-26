import java.util.Arrays;
import java.util.Comparator;

class Solution {
    public int[][] reconstructQueue(int[][] people) {
        Arrays.sort(people, new Comparator<int[]>() {
            @Override
            public int compare(int[] o1, int[] o2) {
                if(o1[0]==o2[0]) {
                    return Integer.compare(o1[1], o2[1]);
                }
                return Integer.compare(o1[0], o2[0]);
            }
        });

        int[][]res = new int[people.length][2];
        for(int[] r:res) {
            r[0] = -1;
        }

        for(int[] p:people) {
            int idx = p[1];
            for(int i=0;i<=idx;i++) {
                if(res[i][0] != -1 && res[i][0] < p[0]) {
                    idx++;
                }
            }
            res[idx]=p;
        }
        return res;
    }
}