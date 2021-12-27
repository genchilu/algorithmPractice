import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    public void TestDev1(){
        var sol = new Solution();
        int[][] heights = new int[][]{
                {1,2,2,3,5},{3,2,3,4,4},{2,4,5,3,1},{6,7,1,4,5},{5,1,1,2,4}
        };

        sol.pacificAtlantic(heights);
    }
}
