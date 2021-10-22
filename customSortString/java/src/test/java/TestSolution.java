import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

public class TestSolution {
    Solution solution = new Solution();

    @Test
    public void TestDev1(){
        String order = "cba";
        String s = "abcd";
        String expect = "cbad";

        String actual = solution.customSortString(order, s);
        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev2(){
        String order = "cbafg";
        String s = "abcd";
        String expect = "cbad";

        String actual = solution.customSortString(order, s);
        Assertions.assertEquals(expect, actual);
    }
}
