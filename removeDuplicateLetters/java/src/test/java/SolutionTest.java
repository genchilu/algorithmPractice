import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

public class SolutionTest {

    @Test
    public void TestDev1(){
        String s = "bcabc";
        String expect = "abc";

        var sol = new Solution();

        var actual = sol.removeDuplicateLetters(s);

        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev2(){
        String s = "cbacdcbc";
        String expect = "acdb";

        var sol = new Solution();

        var actual = sol.removeDuplicateLetters(s);

        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev3(){
        String s = "bbcaac";
        String expect = "bac";

        var sol = new Solution();

        var actual = sol.removeDuplicateLetters(s);

        Assertions.assertEquals(expect, actual);
    }
}
