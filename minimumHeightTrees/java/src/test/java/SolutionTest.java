import org.junit.jupiter.api.Test;

public class SolutionTest {

    @Test
    public void TestDev(){
        var sol = new Solution();
        int n =4;
        int[][] edges = new int[][]{
                {1,0},{1,2},{1,3}
        };

        sol.findMinHeightTrees(n, edges);
    }
}
