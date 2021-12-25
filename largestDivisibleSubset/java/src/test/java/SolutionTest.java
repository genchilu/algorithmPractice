import org.junit.jupiter.api.Test;

public class SolutionTest {

    @Test
    public void TestDev1(){
        int[] nums = new int[]{5,9,18,54,108,540,90,180,360,720};
        var sol = new Solution();
        sol.largestDivisibleSubset(nums);
    }
}
