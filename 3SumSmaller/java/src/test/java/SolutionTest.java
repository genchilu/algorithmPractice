import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    public void TestDev1(){
        var sol = new Solution();
        var nums = new int[]{-2,0,1,3};
        var target = 2;
        var expect = 2;

        var actual = sol.threeSumSmaller(nums,target);

        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev2(){
        var sol = new Solution();
        var nums = new int[]{};
        var target = 0;
        var expect = 0;

        var actual = sol.threeSumSmaller(nums,target);

        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev3(){
        var sol = new Solution();
        var nums = new int[]{0};
        var target = 0;
        var expect = 0;

        var actual = sol.threeSumSmaller(nums,target);

        Assertions.assertEquals(expect, actual);
    }
}
