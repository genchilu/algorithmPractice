import junitparams.JUnitParamsRunner;
import junitparams.Parameters;
import org.junit.Test;
import org.junit.runner.RunWith;

import java.util.List;
import java.util.Map;
import static org.junit.Assert.*;

@RunWith(JUnitParamsRunner.class)
public class KosarajuTest {
    @Test
    @Parameters
    public void testDoKosarajuScc(int[][] inputEdges, int[][] expectComponments) {

        Kosaraju kosaraju = new Kosaraju();
        Map<Integer, List<Integer>> actualComponments = kosaraju.doKosarajuScc(inputEdges);

        assertEquals(expectComponments.length, actualComponments.size());
        for (int[] expectComponment: expectComponments) {
            int lowestVertex = expectComponment[0];
            List<Integer> componment = actualComponments.get(new Integer(lowestVertex));

            assertEquals(expectComponment.length, componment.size());
            for(int i = 0; i < expectComponment.length;i++) {
                assertEquals(expectComponment[i], componment.get(i).intValue());
            }
        }
    }
    private Object[] parametersForTestDoKosarajuScc() {
        return new Object[]{
                new Object[]{
                        null,
                        new int[][]{
                        }
                },
                new Object[]{
                        new int[][]{
                        },
                        new int[][]{
                        }
                },
                new Object[]{
                        new int[][]{
                                new int[]{6,7},
                                new int[]{7,0},
                                new int[]{7,4},
                                new int[]{7,5},
                                new int[]{5,6},
                                new int[]{5,4},
                                new int[]{4,3},
                                new int[]{3,2},
                                new int[]{3,4},
                                new int[]{0,1},
                                new int[]{0,2},
                                new int[]{0,3},
                                new int[]{1,0},
                                new int[]{1,2},
                        },
                        new int[][]{
                                new int[]{5,6,7},
                                new int[]{3,4},
                                new int[]{2},
                                new int[]{0,1},
                        }
                },
                new Object[]{
                        new int[][]{
                                new int[]{1,0},
                                new int[]{2,1},
                                new int[]{0,2},
                                new int[]{0,3},
                                new int[]{3,4},
                        },
                        new int[][]{
                                new int[]{0,1,2},
                                new int[]{3},
                                new int[]{4},
                        }
                },
                new Object[]{
                        new int[][]{
                                new int[]{1,3},
                                new int[]{2,1},
                                new int[]{3,2},
                                new int[]{3,4},
                                new int[]{4,5},
                                new int[]{5,6},
                                new int[]{6,4},
                                new int[]{7,8},
                                new int[]{8,10},
                                new int[]{9,7},
                                new int[]{9,6},
                                new int[]{10,9},
                                new int[]{10,11},
                        },
                        new int[][]{
                                new int[]{1,2,3},
                                new int[]{4,5,6},
                                new int[]{7,8,9,10},
                                new int[]{11},
                        }
                },
        };
    }


}
