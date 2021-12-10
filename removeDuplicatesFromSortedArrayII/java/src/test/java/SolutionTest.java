import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

public class SolutionTest {

    @Test
    public void TestDev1(){
        //removeDuplicates(int[] nums) {

        Solution solution = new Solution();
        int[] nums = new int[]{1,1,1,2,2,3};

        int expectK = 5;
        int[] expectNums = new int[]{1,1,2,2,3};

        var actualk = solution.removeDuplicates(nums);

        Assertions.assertEquals(expectK, actualk);
        for(int i=0;i<expectNums.length;i++){
            Assertions.assertEquals(expectNums[i], nums[i]);
        }
    }

    @Test
    public void TestDev2(){
        //removeDuplicates(int[] nums) {

        Solution solution = new Solution();
        int[] nums = new int[]{0,0,1,1,1,1,2,3,3};

        int expectK = 7;
        int[] expectNums = new int[]{0,0,1,1,2,3,3};

        var actualk = solution.removeDuplicates(nums);

        Assertions.assertEquals(expectK, actualk);
        for(int i=0;i<expectNums.length;i++){
            Assertions.assertEquals(expectNums[i], nums[i]);
        }
    }

    @Test
    public void TestDev3(){
        //removeDuplicates(int[] nums) {

        Solution solution = new Solution();
        int[] nums = new int[]{1,1,1};

        int expectK = 2;
        int[] expectNums = new int[]{1,1};

        var actualk = solution.removeDuplicates(nums);

        Assertions.assertEquals(expectK, actualk);
        for(int i=0;i<expectNums.length;i++){
            Assertions.assertEquals(expectNums[i], nums[i]);
        }
    }
}
