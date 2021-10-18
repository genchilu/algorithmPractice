import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

public class TestSolution {

    Solution solution = new Solution();

    @Test
    public void TestDev() {
        int[] nums = new int[]{3,5,2,3};
        int expect = 7;

        int actual = solution.minPairSum(nums);

        Assertions.assertEquals(expect, actual);

    }

    @Test
    public void TestDev2() {
        int[] nums = new int[]{3,5,4,2,4,6};
        int expect = 8;

        int actual = solution.minPairSum(nums);

        Assertions.assertEquals(expect, actual);

    }

}
