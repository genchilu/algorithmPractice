import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

class Solution {



    public List<List<Integer>> pacificAtlantic(int[][] heights) {
        List<List<Integer>> res = new ArrayList<>();
        boolean[][] pacificFlow = new boolean[heights.length][heights[0].length];
        for(int i =0;i<heights.length;i++) {
            checkFlow(heights,i,0,pacificFlow);
        }
        for(int j =0;j<heights[0].length;j++) {
            checkFlow(heights,0,j,pacificFlow);
        }

        boolean[][] atlanticFlow = new boolean[heights.length][heights[0].length];
        for(int i =0;i<heights.length;i++) {
            checkFlow(heights,i, heights[0].length-1,atlanticFlow);
        }
        for(int j =0;j<heights[0].length;j++) {
            checkFlow(heights,heights.length-1,j,atlanticFlow);
        }

        for(int i =0;i<heights.length;i++){
            for(int j=0;j<heights[0].length;j++){
                if(pacificFlow[i][j] && atlanticFlow[i][j]) {
                    res.add(Arrays.asList(i,j));
                }
            }
        }

        return res;
    }

    public void checkFlow(int[][] heights,int i,int j, boolean[][] flowed) {
        flowed[i][j] = true;

        if(i>0) {
            if(!flowed[i-1][j] && heights[i-1][j]>=heights[i][j]) {
                checkFlow(heights, i-1, j, flowed);
            }
        }

        if(i<heights.length-1) {
            if(!flowed[i+1][j] && heights[i+1][j]>=heights[i][j]) {
                checkFlow(heights, i+1, j, flowed);
            }
        }

        if(j>0) {
            if(!flowed[i][j-1] && heights[i][j-1]>=heights[i][j]) {
                checkFlow(heights, i, j-1, flowed);
            }
        }

        if(j<heights[0].length-1) {
            if(!flowed[i][j+1] && heights[i][j+1]>=heights[i][j]) {
                checkFlow(heights, i, j+1, flowed);
            }
        }

    }
}