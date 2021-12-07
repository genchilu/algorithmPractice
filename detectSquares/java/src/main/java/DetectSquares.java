import java.util.*;

class DetectSquares {

    int[][] coordinate = new int[1001][1001];
    List<Integer[]> points = new LinkedList<>();
    public DetectSquares() {

    }

    public void add(int[] point) {
        var x = point[0];
        var y = point[1];

        coordinate[x][y]++;

        var p = new Integer[2];
        p[0]=x;
        p[1]=y;

        points.add(p);
    }

    public int count(int[] point) {
        var x1 = point[0];
        var y1 = point[1];

        var count = Integer.valueOf(0);
        for (var p2:points) {
            var x2 = p2[0];
            var y2 = p2[1];

            if (x1==x2 || y1==y2 || (Math.abs(x1-x2) != Math.abs(y1-y2))) {
                continue;
            }

            count += coordinate[x1][y2] * coordinate[x2][y1];

        }

        return count;
    }
}