import java.util.Arrays;

class Solution {
    int[][] maze;
    int[][] records;
    public int shortestDistance(int[][] maze, int[] start, int[] destination) {

        this.maze = maze;
        this.records = new int[maze.length][maze[0].length];
        for(int i=0;i<records.length;i++){
            Arrays.fill(records[i], Integer.MAX_VALUE);
        }


        int min = Integer.MAX_VALUE;
        dfsMove(start, 'u', records, 0);
        dfsMove(start, 'd', records, 0);
        dfsMove(start, 'l', records, 0);
        dfsMove(start, 'r', records, 0);

        if(records[destination[0]][destination[1]] != Integer.MAX_VALUE) {
            return records[destination[0]][destination[1]];
        }
        return -1;
    }

    public void dfsMove(int[] cur, char dir, int[][] records, int preCount) {


        int count = 0;
        int i = cur[0];
        int j = cur[1];
        switch (dir) {
            case 'u':
                while(i-1 >= 0 && maze[i-1][j] == 0) {
                    i--;
                    count++;
                }
                break;
            case 'd':
                while(i+1 < maze.length && maze[i+1][j] == 0) {
                    i++;
                    count++;
                }
                break;
            case 'l':
                while(j-1>=0 && maze[i][j-1] == 0) {
                    j--;
                    count++;
                }
                break;
            case 'r':
                while(j+1<maze[0].length && maze[i][j+1] == 0){
                    j++;
                    count++;
                }
                break;
            default:
                break;
        }

        int totalCount = preCount+count;
        if(totalCount<records[i][j]) {
            records[i][j] = totalCount;

            dfsMove(new int[]{i,j}, 'u', records, totalCount);
            dfsMove(new int[]{i,j}, 'd', records, totalCount);
            dfsMove(new int[]{i,j}, 'l', records, totalCount);
            dfsMove(new int[]{i,j}, 'r', records, totalCount);
        }

    }
}