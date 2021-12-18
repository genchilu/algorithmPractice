import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    public void TestDev1(){
        Solution sol = new Solution();
        String s = "";
        String t = "";

        boolean expect = false;

        boolean actual = sol.isOneEditDistance(s, t);

        Assertions.assertEquals(expect, actual);

    }

    @Test
    public void TestDev2(){
        Solution sol = new Solution();
        String s = "a";
        String t = "";

        boolean expect = true;

        boolean actual = sol.isOneEditDistance(s, t);

        Assertions.assertEquals(expect, actual);

    }

    @Test
    public void TestDev3(){
        Solution sol = new Solution();
        String s = "a";
        String t = "bb";

        boolean expect = false;

        boolean actual = sol.isOneEditDistance(s, t);

        Assertions.assertEquals(expect, actual);

    }

    @Test
    public void TestDev4(){
        Solution sol = new Solution();
        String s = "bb";
        String t = "bb";

        boolean expect = false;

        boolean actual = sol.isOneEditDistance(s, t);

        Assertions.assertEquals(expect, actual);

    }

    @Test
    public void TestDev5(){
        Solution sol = new Solution();
        String s = "bab";
        String t = "bb";

        boolean expect = true;

        boolean actual = sol.isOneEditDistance(s, t);

        Assertions.assertEquals(expect, actual);

    }

    @Test
    public void TestDev6(){
        Solution sol = new Solution();
        String s = "bba";
        String t = "bb";

        boolean expect = true;

        boolean actual = sol.isOneEditDistance(s, t);

        Assertions.assertEquals(expect, actual);

    }

    @Test
    public void TestDev7(){
        Solution sol = new Solution();
        String s = "abb";
        String t = "bb";

        boolean expect = true;

        boolean actual = sol.isOneEditDistance(s, t);

        Assertions.assertEquals(expect, actual);

    }

    @Test
    public void TestDev8(){
        Solution sol = new Solution();
        String s = "abb";
        String t = "bba";

        boolean expect = false;

        boolean actual = sol.isOneEditDistance(s, t);

        Assertions.assertEquals(expect, actual);

    }
}
