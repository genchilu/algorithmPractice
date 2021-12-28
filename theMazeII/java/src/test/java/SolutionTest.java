import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    public void TestDev1(){
        var sol = new Solution();
        int[][] maze = new int[][]{
                {0,0,1,0,0},{0,0,0,0,0},{0,0,0,1,0},{1,1,0,1,1},{0,0,0,0,0}
        };
        int[] start = new int[]{0,4};
        int[] destination = new int[]{4,4};
        int expect = 12;

        int actual = sol.shortestDistance(maze, start, destination);

        Assertions.assertEquals(expect, actual);
    }
}
