import java.util.ArrayDeque;
import java.util.List;
import java.util.Queue;

class Solution {
    public int[][] updateMatrix(int[][] mat) {

        int[][] res = new int[mat.length][mat[0].length];
        Queue<Integer[]> q = new ArrayDeque<>();
        for(int i=0;i<mat.length;i++){
            for(int j=0;j<mat[0].length;j++){
                if(mat[i][j]==0) {
                    res[i][j] = 0;
                    q.add(new Integer[]{i,j});
                } else {
                    res[i][j] = -1;
                }
            }
        }

        while(q.size()>0){
            Integer[] point = q.remove();
            int i = point[0];
            int j = point[1];

            if(i >0 && res[i-1][j] == -1) {
                res[i-1][j] = res[i][j]+1;
                q.add(new Integer[]{i-1,j});
            }
            if(i <res.length-1 && res[i+1][j] == -1) {
                res[i+1][j] = res[i][j]+1;
                q.add(new Integer[]{i+1,j});
            }
            if(j >0 && res[i][j-1] == -1) {
                res[i][j-1] = res[i][j]+1;
                q.add(new Integer[]{i,j-1});
            }
            if(j <res[0].length-1 && res[i][j+1] == -1) {
                res[i][j+1] = res[i][j]+1;
                q.add(new Integer[]{i,j+1});
            }
        }

        return res;
    }
}