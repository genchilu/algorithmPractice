import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    public void TestDev1(){
        int[][] points = new int[][]{
                {0,0},{1,0},{2,0}
        };
        int expect = 2;
        Solution sol = new Solution();

        int actual = sol.numberOfBoomerangs(points);

        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev2(){
        int[][] points = new int[][]{
                {1,1},{2,2},{3,3}
        };
        int expect = 2;
        Solution sol = new Solution();

        int actual = sol.numberOfBoomerangs(points);

        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev3(){
        int[][] points = new int[][]{
                {1,1}
        };
        int expect = 0;
        Solution sol = new Solution();

        int actual = sol.numberOfBoomerangs(points);

        Assertions.assertEquals(expect, actual);
    }
}
