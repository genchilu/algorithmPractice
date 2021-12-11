import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

public class SolutionTest {

    @Test
    public void TestDev1(){
        var s = new Solution();
        //findPivot(int[] nums) {
        int[] nums = new int[]{4,5,6,6,7,0,1,2,4,4};
        int expect = 5;

        var actual = s.findPivot(nums);

        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev2(){
        var s = new Solution();
        //findPivot(int[] nums) {
        int[] nums = new int[]{0,1,2,4,4,4,5,6,6,7};
        int expect = 0;

        var actual = s.findPivot(nums);

        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev3(){
        var s = new Solution();
        //findPivot(int[] nums) {
        int[] nums = new int[]{7,0,1,2,4,4,4,5,6,6};
        int expect = 1;

        var actual = s.findPivot(nums);

        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev4(){
        var s = new Solution();
        //findPivot(int[] nums) {
        int[] nums = new int[]{7};
        int expect = 0;

        var actual = s.findPivot(nums);

        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev5(){
        var s = new Solution();
        //public boolean search(int[] nums, int target) {
        int[] nums = new int[]{2,5,6,0,0,1,2};
        int target = 0;
        boolean expect = true;

        var actual = s.search(nums,target);

        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev6(){
        var s = new Solution();
        //public boolean search(int[] nums, int target) {
        int[] nums = new int[]{2,5,6,0,0,1,2};
        int target = 3;
        boolean expect = false;

        var actual = s.search(nums,target);

        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev7(){
        var s = new Solution();
        //findPivot(int[] nums) {
        int[] nums = new int[]{1,0,1,1};
        int expect = 1;

        var actual = s.findPivot(nums);

        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev8(){
        var s = new Solution();
        //public boolean search(int[] nums, int target) {
        int[] nums = new int[]{1,0,1,1};
        int target = 0;
        boolean expect = true;

        var actual = s.search(nums,target);

        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev9(){
        var s = new Solution();
        //public boolean search(int[] nums, int target) {
        int[] nums = new int[]{1,0,1,1,1};
        int target = 0;
        boolean expect = true;

        var actual = s.search(nums,target);

        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev10(){
        var s = new Solution();
        //findPivot(int[] nums) {
        int[] nums = new int[]{1,0,1,1,1};
        int expect = 1;

        var actual = s.findPivot(nums);

        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev11(){
        var s = new Solution();
        //findPivot(int[] nums) {
        int[] nums = new int[]{1,3};
        int target = 0;
        boolean expect = false;

        var actual = s.search(nums,target);

        Assertions.assertEquals(expect, actual);
    }

}
