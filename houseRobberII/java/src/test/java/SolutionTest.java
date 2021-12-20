import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

public class SolutionTest {

    @Test
    public void TestDev1(){
        var sol = new Solution();
        int[] nums = new int[]{2,3,2};
        int expect = 3;

        int actual = sol.rob(nums);

        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev2(){
        var sol = new Solution();
        int[] nums = new int[]{1,2,3,1};
        int expect = 4;

        int actual = sol.rob(nums);

        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev4(){
        var sol = new Solution();
        int[] nums = new int[]{1,2,3,};
        int expect = 3;

        int actual = sol.rob(nums);

        Assertions.assertEquals(expect, actual);
    }
}
