import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    public void TestDev1(){
        int n = 11;
        int expect = 0;

        var sol = new Solution();

        var acutl = sol.findNthDigit(n);

        Assertions.assertEquals(expect, acutl);
    }

    @Test
    public void TestDev2(){
        int n = 10;
        int expect = 1;

        var sol = new Solution();

        var acutl = sol.findNthDigit(n);

        Assertions.assertEquals(expect, acutl);
    }

    @Test
    public void TestDev3(){
        int n = 12;
        int expect = 1;

        var sol = new Solution();

        var acutl = sol.findNthDigit(n);

        Assertions.assertEquals(expect, acutl);
    }

    @Test
    public void TestDev4(){
        int n = 13;
        int expect = 1;

        var sol = new Solution();

        var acutl = sol.findNthDigit(n);

        Assertions.assertEquals(expect, acutl);
    }

    @Test
    public void TestDev5(){
        int n = 190;
        int expect = 1;

        var sol = new Solution();

        var acutl = sol.findNthDigit(n);

        Assertions.assertEquals(expect, acutl);
    }

    @Test
    public void TestDev6(){
        int n = 191;
        int expect = 0;

        var sol = new Solution();

        var acutl = sol.findNthDigit(n);

        Assertions.assertEquals(expect, acutl);
    }

    @Test
    public void TestDev7(){
        int n = 192;
        int expect = 0;

        var sol = new Solution();

        var acutl = sol.findNthDigit(n);

        Assertions.assertEquals(expect, acutl);
    }

    @Test
    public void TestDev8(){
        int n = 196;
        int expect = 1;

        var sol = new Solution();

        var acutl = sol.findNthDigit(n);

        Assertions.assertEquals(expect, acutl);
    }


    @Test
    public void TestDev9(){
        int n = 197;
        int expect = 0;

        var sol = new Solution();

        var acutl = sol.findNthDigit(n);

        Assertions.assertEquals(expect, acutl);
    }

    @Test
    public void TestDev10(){
        int n = 198;
        int expect = 2;

        var sol = new Solution();

        var acutl = sol.findNthDigit(n);

        Assertions.assertEquals(expect, acutl);
    }

    @Test
    public void TestDev11(){
        int n = 1000000000;
        int expect = 1;

        var sol = new Solution();

        var acutl = sol.findNthDigit(n);

        Assertions.assertEquals(expect, acutl);
    }
}
