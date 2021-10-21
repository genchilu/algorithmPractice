import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

public class TestSolution {
    Solution solution = new Solution();

    @Test
    public void TestDev1(){
        int[] doubled = new int[]{1,3,4,2,6,8};
        int[] expect = new int[]{1,3,4};

        int[] actual = solution.findOriginalArray(doubled);

        Assertions.assertArrayEquals(expect, actual);
    }

    @Test
    public void TestDev2(){
        int[] doubled = new int[]{};
        int[] expect = new int[]{};

        int[] actual = solution.findOriginalArray(doubled);

        Assertions.assertArrayEquals(expect, actual);
    }

    @Test
    public void TestDev3(){
        int[] doubled = new int[]{1};
        int[] expect = new int[]{};

        int[] actual = solution.findOriginalArray(doubled);

        Assertions.assertArrayEquals(expect, actual);
    }
}
