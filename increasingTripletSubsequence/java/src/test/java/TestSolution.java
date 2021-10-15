import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

public class TestSolution {
    private Solution solution = new Solution();

    @Test
    public void TestDev1() {
        int[] nums = new int[]{1,2,3,4,5};
        var actual = solution.increasingTriplet(nums);
        Assertions.assertEquals(true, actual);
    }

    @Test
    public void TestDev2() {
        int[] nums = new int[]{5,4,3,2,1};
        var actual = solution.increasingTriplet(nums);
        Assertions.assertEquals(false, actual);
    }

    @Test
    public void TestDev3() {
        int[] nums = new int[]{2,1,5,0,4,6};
        var actual = solution.increasingTriplet(nums);
        Assertions.assertEquals(true, actual);
    }

}
