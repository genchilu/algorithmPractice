import java.util.ArrayDeque;
import java.util.Queue;

class Solution {
    public void wallsAndGates(int[][] rooms) {
        Queue<Integer[]> queue = new ArrayDeque<>();
        var m = rooms.length;
        var n = rooms.length;

        for(int i=0;i<m;i++){
            for (int j=0;j<n;j++){
                if(rooms[i][j] == 0) {
                    var grid = new Integer[2];
                    grid[0] = i;
                    grid[1] = j;
                    queue.add(grid);
                }
            }
        }

        while(queue.size()>0) {
            Integer[] grid = queue.remove();
            var i = grid[0];
            var j = grid[1];
            //down
            if(canWalk(rooms, i+1, j)) {
                rooms[i+1][j] = rooms[i][j]+1;
                queue.add(new Integer[]{i+1,j});
            }
            //up
            if(canWalk(rooms, i-1, j)) {
                rooms[i-1][j] = rooms[i][j]+1;
                queue.add(new Integer[]{i-1,j});
            }
            //left
            if(canWalk(rooms, i, j-1)) {
                rooms[i][j-1] = rooms[i][j]+1;
                queue.add(new Integer[]{i,j-1});
            }
            //right
            if(canWalk(rooms, i, j+1)) {
                rooms[i][j+1] = rooms[i][j]+1;
                queue.add(new Integer[]{i,j+1});
            }
        }
    }

    public boolean canWalk(int[][] rooms, int i, int j) {
        if(i<0 || i >= rooms.length || j <0 || j>=rooms[0].length) {
            return false;
        }

        if(rooms[i][j] != 2147483647) {
            System.out.format("%d, %d\n", i, j);
            return false;
        }

        return true;
    }
}