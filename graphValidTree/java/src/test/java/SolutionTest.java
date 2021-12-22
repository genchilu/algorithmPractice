import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

public class SolutionTest {

    @Test
    public void TestDev1(){
        var sol = new Solution();
        var edges = new int[][]{
                {0,1},{0,2},{0,3},{1,4}
        };
        var n = 5;

        boolean expect = true;

        var actual = sol.validTree(n, edges);

        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev2(){
        var sol = new Solution();
        var edges = new int[][]{
                {0,1}
        };
        var n = 5;

        boolean expect = false;

        var actual = sol.validTree(n, edges);

        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev3(){
        var sol = new Solution();
        var edges = new int[][]{
                {0,1},{1,2},{2,3},{1,3},{1,4}
        };
        var n = 5;

        boolean expect = false;

        var actual = sol.validTree(n, edges);

        Assertions.assertEquals(expect, actual);
    }
}
