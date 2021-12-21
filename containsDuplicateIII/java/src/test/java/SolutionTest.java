import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    public void TestDev1(){
        // containsNearbyAlmostDuplicate(int[] nums, int k, int t) {
        var sol = new Solution();
        int[] nums = new int[]{1,2,3,1};
        int k = 3;
        int t = 0;
        boolean expect = true;

        boolean actual = sol.containsNearbyAlmostDuplicate(nums,k,t);

        Assertions.assertEquals(expect, actual);

    }

    @Test
    public void TestDev2(){
        // containsNearbyAlmostDuplicate(int[] nums, int k, int t) {
        var sol = new Solution();
        int[] nums = new int[]{1,0,1,1};
        int k = 1;
        int t = 2;
        boolean expect = true;

        boolean actual = sol.containsNearbyAlmostDuplicate(nums,k,t);

        Assertions.assertEquals(expect, actual);

    }

    @Test
    public void TestDev3(){
        // containsNearbyAlmostDuplicate(int[] nums, int k, int t) {
        var sol = new Solution();
        int[] nums = new int[]{1,5,9,1,5,9};
        int k = 2;
        int t = 3;
        boolean expect = false;

        boolean actual = sol.containsNearbyAlmostDuplicate(nums,k,t);

        Assertions.assertEquals(expect, actual);

    }

    @Test
    public void TestDev4(){
        // containsNearbyAlmostDuplicate(int[] nums, int k, int t) {
        var sol = new Solution();
        int[] nums = new int[]{1,5,9,14,19,24};
        int k = 100;
        int t = 3;
        boolean expect = false;

        boolean actual = sol.containsNearbyAlmostDuplicate(nums,k,t);

        Assertions.assertEquals(expect, actual);

    }

    @Test
    public void TestDev5(){
        // containsNearbyAlmostDuplicate(int[] nums, int k, int t) {
        var sol = new Solution();
        int[] nums = new int[]{2,5,7,1,5,9};
        int k = 100;
        int t = 3;
        boolean expect = true;

        boolean actual = sol.containsNearbyAlmostDuplicate(nums,k,t);

        Assertions.assertEquals(expect, actual);

    }

    @Test
    public void TestDev6(){
        // containsNearbyAlmostDuplicate(int[] nums, int k, int t) {
        var sol = new Solution();
        int[] nums = new int[]{-2147483648,2147483647};
        int k = 1;
        int t = 1;
        boolean expect = false;

        boolean actual = sol.containsNearbyAlmostDuplicate(nums,k,t);

        Assertions.assertEquals(expect, actual);

    }

    @Test
    public void TestDev7(){
        // containsNearbyAlmostDuplicate(int[] nums, int k, int t) {
        var sol = new Solution();
        int[] nums = new int[]{2147483647,-1,2147483647};
        int k = 1;
        int t = 2147483647;
        boolean expect = false;

        boolean actual = sol.containsNearbyAlmostDuplicate(nums,k,t);

        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev8(){
        // containsNearbyAlmostDuplicate(int[] nums, int k, int t) {
        var sol = new Solution();
        int[] nums = new int[]{1,2};
        int k = 0;
        int t = 1;
        boolean expect = false;

        boolean actual = sol.containsNearbyAlmostDuplicate(nums,k,t);

        Assertions.assertEquals(expect, actual);
    }

}
