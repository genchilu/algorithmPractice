import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

public class SolutionTest {

    //public int uniquePathsWithObstacles(int[][] obstacleGrid) {

    @Test
    public void TestDev1() {
        var solution = new Solution();
        var obstacleGrid = new int[][]{
                {0,0,0},
                {0,1,0},
                {0,0,0},
        };
        var expect = 2;

        var actual = solution.uniquePathsWithObstacles(obstacleGrid);

        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev2() {
        var solution = new Solution();
        var obstacleGrid = new int[][]{
                {0,1},
                {0,0},
        };
        var expect = 1;

        var actual = solution.uniquePathsWithObstacles(obstacleGrid);

        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev3() {
        var solution = new Solution();
        var obstacleGrid = new int[][]{
                {0,0},
                {0,1},
        };
        var expect = 0;

        var actual = solution.uniquePathsWithObstacles(obstacleGrid);

        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev4() {
        var solution = new Solution();
        var obstacleGrid = new int[][]{
                {0,1},
        };
        var expect = 0;

        var actual = solution.uniquePathsWithObstacles(obstacleGrid);

        Assertions.assertEquals(expect, actual);
    }
}
