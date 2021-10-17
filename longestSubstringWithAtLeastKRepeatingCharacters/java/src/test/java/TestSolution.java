import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

public class TestSolution {

    Solution solution = new Solution();
    @Test
    public void TestDev(){
        String s = "aaabb";
        int k = 3;
        int expect = 3;

        int actual = solution.longestSubstring(s,k);

        Assertions.assertEquals(expect, actual);

    }

    @Test
    public void TestDev2(){
        String s = "ababbc";
        int k = 2;
        int expect = 5;

        int actual = solution.longestSubstring(s,k);

        Assertions.assertEquals(expect, actual);

    }

    @Test
    public void TestDev3(){
        String s = "aaabbb";
        int k = 3;
        int expect = 6;

        int actual = solution.longestSubstring(s,k);

        Assertions.assertEquals(expect, actual);

    }
}
