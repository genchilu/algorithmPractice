import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    public void TestDev1(){
        var sol = new Solution();
        int[] nums = new int[]{-4,-2,2,4};
        int a=1;
        int b=3;
        int c=5;
        int[] expect = new int[]{3,9,15,33};

        int[] actual = sol.sortTransformedArray(nums, a,b,c);

        Assertions.assertArrayEquals(expect,actual);
    }

    @Test
    public void TestDev2(){
        var sol = new Solution();
        int[] nums = new int[]{-4,-2,2,4};
        int a=-1;
        int b=3;
        int c=5;
        int[] expect = new int[]{-23,-5,1,7};

        int[] actual = sol.sortTransformedArray(nums, a,b,c);

        Assertions.assertArrayEquals(expect,actual);
    }
}
