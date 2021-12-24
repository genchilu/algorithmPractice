import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    public void TestDev1(){
        int[] primes = new int[]{2,7,13,19};
        int n = 12;
        int expect = 32;

        var sol = new Solution();

        var actual = sol.nthSuperUglyNumber(n, primes);

        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev2(){
        int[] primes = new int[]{2,3,5};
        int n = 1;
        int expect = 1;

        var sol = new Solution();

        var actual = sol.nthSuperUglyNumber(n, primes);

        Assertions.assertEquals(expect, actual);
    }
}
