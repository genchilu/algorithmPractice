import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.concurrent.Flow;

class Solution {
    class FlowRes{
        boolean toAtlantic;
        boolean toPacific;

        public FlowRes(boolean toAtlantic, boolean toPacific) {
            this.toAtlantic = toAtlantic;
            this.toPacific = toPacific;
        }
    }


    public List<List<Integer>> pacificAtlantic(int[][] heights) {
        List<List<Integer>> res = new ArrayList<>();
        boolean[][] seen = new boolean[heights.length][heights[0].length];

        //0 ori / 1 up/2 down/3 left/4 right
        FlowRes[][][] cache = new FlowRes[heights.length][heights[0].length][5];

        for(int i = 0;i<heights.length;i++){
            for(int j=0;j<heights[0].length;j++){
                FlowRes flowRes = checkFlow(heights, i, j, seen, 0, cache);
                if(flowRes.toAtlantic && flowRes.toPacific) {
                    res.add(Arrays.asList(i,j));
                }
            }
        }

        return res;
    }

    public FlowRes checkFlow(int[][] heights, int i, int j, boolean[][] seen, int direction, FlowRes[][][] cache) {
        if(cache[i][j][direction] != null) {
            return cache[i][j][direction];
        }

        FlowRes res = new FlowRes(false, false);
        if(i==0 || j ==0) {
            res.toPacific = true;
        }
        if(i==heights.length-1 || j == heights[0].length-1) {
            res.toAtlantic = true;
        }

        seen[i][j] = true;
        if(i>0) {
            if(!seen[i-1][j] && heights[i-1][j] <= heights[i][j]) {
                FlowRes nextRes = checkFlow(heights, i-1,j,seen, 1, cache);
                updateFlowRes(nextRes, res);
            }
        }

        if(i<heights.length-1) {
            if(!seen[i+1][j] && heights[i+1][j]<=heights[i][j]) {
                FlowRes nextRes = checkFlow(heights, i+1,j,seen, 2, cache);
                updateFlowRes(nextRes, res);
            }
        }

        if(j>0) {
            if(!seen[i][j-1] && heights[i][j-1]<=heights[i][j]) {
                FlowRes nextRes = checkFlow(heights, i,j-1,seen, 3, cache);
                updateFlowRes(nextRes, res);
            }
        }

        if(j<heights[0].length-1) {
            if(!seen[i][j+1] && heights[i][j+1]<=heights[i][j]) {
                FlowRes nextRes = checkFlow(heights, i,j+1,seen, 4, cache);
                updateFlowRes(nextRes, res);
            }
        }
        seen[i][j] = false;

        cache[i][j][direction] = res;
        return res;
    }

    public FlowRes updateFlowRes(FlowRes next, FlowRes cur) {
        if(next.toPacific) {
            cur.toPacific = true;
        }

        if(next.toAtlantic) {
            cur.toAtlantic = true;
        }

        return cur;
    }
}