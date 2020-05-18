import junitparams.Parameters;
import org.junit.Test;
import org.junit.runner.RunWith;
import junitparams.JUnitParamsRunner;
import static org.junit.Assert.*;

@RunWith(JUnitParamsRunner.class)
public class CountInversionsTest {

    @Test
    @Parameters
    public void testCountInversions(int[] input, int expect) {
        int actualResult = CountingInversions.CountInversions(input);

        assertEquals(expect, actualResult);

    }
    private Object[] parametersForTestCountInversions() {
        return new Object[]{
                new Object[]{null, 0},
                new Object[]{new int[] {}, 0},
                new Object[]{new int[] {1}, 0},
                new Object[]{new int[] {1,2}, 0},
                new Object[]{new int[] {2,1}, 1},
                new Object[]{new int[] {1,1}, 0},
                new Object[]{new int[] {1,2,3,4,5}, 0},
                new Object[]{new int[] {5,4,3,2,1}, 10},
                new Object[]{new int[] {7,5,3,3,9,6,7}, 8},
                new Object[]{new int[] {23, 42, 32, 64, 12, 4}, 10},
                new Object[]{new int[] {2,1,2,2,2,2,2}, 1},
                new Object[]{new int[] {10,12,9,13,8,14,7}, 12},
        };
    }
}
