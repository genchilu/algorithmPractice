import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    public void TestDev(){
        var nums = new int[]{3,5,2,1,6,4};
        var expect = new int[]{3,5,2,6,1,3};

        var sol = new Solution();

        sol.wiggleSort(nums);

        Assertions.assertEquals(expect, nums);

    }
}
