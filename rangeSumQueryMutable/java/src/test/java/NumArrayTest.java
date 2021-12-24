import org.junit.jupiter.api.Test;

public class NumArrayTest {

    @Test
    public void TestDev(){
        int[] nums = new int[]{0,9,5,7,3};
        NumArray numArray = new NumArray(nums);
        //numArray.sumRange(4,4);
        numArray.sumRange(2,4);
        numArray.sumRange(3,3);
    }
}
