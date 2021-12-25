import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

public class SolutionTest {

    @Test
    public void TestDev1(){
        //minSubArrayLen(int target, int[] nums)
        var sol = new Solution();

        int target = 7;
        int[] nums = new int[]{2,3,1,2,4,3};
        int expect = 2;

        int actual = sol.minSubArrayLen(target, nums);

        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev2(){
        //minSubArrayLen(int target, int[] nums)
        var sol = new Solution();

        int target = 4;
        int[] nums = new int[]{1,4,4};
        int expect = 1;

        int actual = sol.minSubArrayLen(target, nums);

        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev3(){
        //minSubArrayLen(int target, int[] nums)
        var sol = new Solution();

        int target = 11;
        int[] nums = new int[]{1,1,1,1,1,1,1,1};
        int expect = 0;

        int actual = sol.minSubArrayLen(target, nums);

        Assertions.assertEquals(expect, actual);
    }
}
