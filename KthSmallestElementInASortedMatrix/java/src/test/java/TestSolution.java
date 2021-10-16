import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

public class TestSolution {

    Solution solution = new Solution();

    @Test
    public void TestDev1(){
        var matrix = new int[][]{
                {1,5,9},{10,11,13},{12,13,15}
        };

        var actual = solution.kthSmallest(matrix, 8);

        Assertions.assertEquals(13, actual);
    }

    @Test
    public void TestDev2(){
        var matrix = new int[][]{
                {-5}
        };

        var actual = solution.kthSmallest(matrix, 1);

        Assertions.assertEquals(-5, actual);
    }

    @Test
    public void TestDev3(){
        var matrix = new int[][]{
                {1,2},{1,3}
        };

        var actual = solution.kthSmallest(matrix, 2);

        Assertions.assertEquals(1, actual);
    }

    @Test
    public void TestDev4(){
        var matrix = new int[][]{
                {1,4,7,11,15},{2,5,8,12,19},{3,6,9,16,22},{10,13,14,17,24},{18,21,23,26,30}
        };

        var actual = solution.kthSmallest(matrix, 5);

        Assertions.assertEquals(5, actual);
    }
}
