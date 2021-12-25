import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    public void TestDev1(){
        int[] nums = new int[]{1,2,3};
        int target = 4;
        int expect = 7;
        var sol = new Solution();

        var actual = sol.combinationSum4(nums, target);

        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev2(){
        int[] nums = new int[]{9};
        int target = 3;
        int expect = 0;
        var sol = new Solution();

        var actual = sol.combinationSum4(nums, target);

        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev3(){
        int[] nums = new int[]{4,2,1};
        int target = 32;
        int expect = 0;
        var sol = new Solution();

        var actual = sol.combinationSum4(nums, target);

        Assertions.assertEquals(expect, actual);
    }
}
