import junitparams.JUnitParamsRunner;
import junitparams.Parameters;
import org.junit.Test;
import org.junit.runner.RunWith;
import static org.junit.Assert.*;

@RunWith(JUnitParamsRunner.class)
public class KargerMinCutTest {

    @Test
    @Parameters
    public void testDoMinCut(int[][] edges, int expectMinCut){
        int actualMinCut = KargerMinCut.doMinCut(edges);
        assertEquals(expectMinCut, actualMinCut);
    }
    private Object[] parametersForTestDoMinCut() {
        return new Object[]{
                new Object[]{
                        new int[][]{},
                        0,
                },
                new Object[]{
                        null,
                        0,
                },
                new Object[]{
                        new int[][]{
                                new int[]{0,1},
                                new int[]{1,2},
                                new int[]{2,3},
                                new int[]{0,4},
                                new int[]{1,5},
                                new int[]{2,6},
                                new int[]{3,7},
                                new int[]{4,5},
                                new int[]{5,6},
                                new int[]{6,7},
                                new int[]{0,5},
                                new int[]{1,4},
                                new int[]{2,7},
                                new int[]{3,6},
                        },
                        2,
                },
                new Object[]{
                        new int[][]{
                                new int[]{0,1},
                                new int[]{0,2},
                                new int[]{0,3},
                                new int[]{0,4},
                                new int[]{1,2},
                                new int[]{1,3},
                                new int[]{1,4},
                                new int[]{2,3},
                                new int[]{2,4},
                                new int[]{3,4},
                                new int[]{5,6},
                                new int[]{5,7},
                                new int[]{5,8},
                                new int[]{5,9},
                                new int[]{6,7},
                                new int[]{6,8},
                                new int[]{6,9},
                                new int[]{7,8},
                                new int[]{7,9},
                                new int[]{8,9},
                                new int[]{0,5},
                                new int[]{2,6},
                                new int[]{4,7},
                        },
                        3,
                },
        };
    }
}
