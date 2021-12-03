import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

public class SolutionTest {

    @Test
    public void TestDev1(){
        var solution = new Solution();
        var n = 10;
        var expect = 12;

        var actual = solution.nthUglyNumber(n);

        Assertions.assertEquals(expect,actual);
    }

    @Test
    public void TestDev2(){
        var solution = new Solution();
        var n = 1;
        var expect = 1;

        var actual = solution.nthUglyNumber(n);

        Assertions.assertEquals(expect,actual);
    }

    @Test
    public void TestDev3(){
        var solution = new Solution();
        var n = 11;
        var expect = 15;

        var actual = solution.nthUglyNumber(n);

        Assertions.assertEquals(expect,actual);
    }
}
