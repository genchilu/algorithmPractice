import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;



public class TestSolution {
    Solution solution = new Solution();

    @Test
    public void TestDev1(){
        int[] nums = new int[]{1,5,11,5};
        boolean expect = true;

        boolean actual = solution.canPartition(nums);

        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev2(){
        int[] nums = new int[]{1,2,3,5};
        boolean expect = false;

        boolean actual = solution.canPartition(nums);

        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev3(){
        int[] nums = new int[]{100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,99,97};
        boolean expect = false;

        boolean actual = solution.canPartition(nums);

        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev4(){
        int[] nums = new int[]{1,1,1,1};
        boolean expect = true;

        boolean actual = solution.canPartition(nums);

        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev5(){
        int[] nums = new int[]{2,2,1,1};
        boolean expect = true;

        boolean actual = solution.canPartition(nums);

        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev6(){
        int[] nums = new int[]{1,2,3,4,5,6,7};
        boolean expect = true;

        boolean actual = solution.canPartition(nums);

        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev7(){
        int[] nums = new int[]{97,100,88,49,43,55,2,62,72,13,60,36,67,64,13,39,66,6,26,45,46,75,28,66,71,75,91,33,64,54,3,76,76,35,49,6,63,11,80,86,73,95,60,58,61,42,52,27,73,51,58,38,28,62,84,55,90,40,52,96,55,32,52,63,44,90,14,68,50,32,73,64,92,74,66,64,22,51,27,30,83,30,37,25,22,46,95,59,57,21,29,72,7,56,47,48,54,75,67,33};
        boolean expect = true;

        boolean actual = solution.canPartition(nums);

        Assertions.assertEquals(expect, actual);
    }

}
