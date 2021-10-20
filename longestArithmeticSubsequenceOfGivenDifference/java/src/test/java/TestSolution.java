import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

public class TestSolution {

    Solution solution = new Solution();

    @Test
    public void TestDev1(){
        int[] arr = new int[]{1,2,3,4};
        int difference = 1;
        int expect = 4;

        int actual = solution.longestSubsequence(arr, difference);

        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev2(){
        int[] arr = new int[]{1,3,5,7};
        int difference = 1;
        int expect = 1;

        int actual = solution.longestSubsequence(arr, difference);

        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev3(){
        int[] arr = new int[]{1,5,7,8,5,3,4,2,1};
        int difference = -2;
        int expect = 4;

        int actual = solution.longestSubsequence(arr, difference);

        Assertions.assertEquals(expect, actual);
    }
}
