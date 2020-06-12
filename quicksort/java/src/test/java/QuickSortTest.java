import junitparams.JUnitParamsRunner;
import junitparams.Parameters;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.junit.runners.Parameterized;



import static org.junit.Assert.*;

@RunWith(JUnitParamsRunner.class)
public class QuickSortTest {

    @Test
    @Parameters
    public void testSort(int[] input, int[] expect){
        QuickSort.sort(input);
        assertArrayEquals(expect, input);
    }
    private Object[] parametersForTestSort() {
        return new Object[]{
                new Object[]{null, null},
                new Object[]{new int[] {}, new int[] {}},
                new Object[]{new int[] {1}, new int[] {1}},
                new Object[]{new int[] {1,2}, new int[] {1,2}},
                new Object[]{new int[] {2,1}, new int[] {1,2}},
                new Object[]{new int[] {1,1}, new int[] {1,1}},
                new Object[]{new int[] {1,2,3,4,5}, new int[] {1,2,3,4,5}},
                new Object[]{new int[] {5,4,3,2,1}, new int[] {1,2,3,4,5}},
                new Object[]{new int[] {7,5,3,3,9,6,7}, new int[] {3,3,5,6,7,7,9}},
                new Object[]{new int[] {2,1,2,2,2,2,2}, new int[] {1,2,2,2,2,2,2}},
                new Object[]{new int[] {10,12,9,13,8,14,7}, new int[] {7,8,9,10,12,13,14}},
        };
    }
}
