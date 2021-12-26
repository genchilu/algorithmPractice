import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    public void TestDev1(){
        int[][] people = new int[][]{
                {7,0},{4,4},{7,1},{5,0},{6,1},{5,2}
        };

        var sol = new Solution();
        var expect = new int[][]{
                {5,0},{7,0},{5,2},{6,1},{4,4},{7,1}
        };

        var actual = sol.reconstructQueue(people);

        Assertions.assertArrayEquals(expect, actual);
    }
    @Test
    public void TestDev2(){
        int[][] people = new int[][]{
                {6,0},{5,0},{4,0},{3,2},{2,2},{1,4}
        };

        var sol = new Solution();
        var expect = new int[][]{
                {4,0},{5,0},{2,2},{3,2},{1,4},{6,0}
        };

        var actual = sol.reconstructQueue(people);

        Assertions.assertArrayEquals(expect, actual);
    }

}
