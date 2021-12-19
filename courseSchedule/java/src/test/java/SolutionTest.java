import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    public void TestDev1(){
        Solution sol = new Solution();

        int numCourses = 2;
        int[][] prerequisites = new int[][]{
                {1,0},
        };

        boolean expect = true;

        boolean actual = sol.canFinish(numCourses, prerequisites);

        Assertions.assertEquals(expect, actual);

    }

    @Test
    public void TestDev2(){
        Solution sol = new Solution();

        int numCourses = 2;
        int[][] prerequisites = new int[][]{
                {1,0},
                {0,1},
        };

        boolean expect = false;

        boolean actual = sol.canFinish(numCourses, prerequisites);

        Assertions.assertEquals(expect, actual);

    }

    @Test
    public void TestDev3(){
        Solution sol = new Solution();


        int numCourses = 4;
        int[][] prerequisites = new int[][]{
                {0,1},
                {3,1},
                {1,3},
                {3,2}
        };

        boolean expect = false;

        boolean actual = sol.canFinish(numCourses, prerequisites);

        Assertions.assertEquals(expect, actual);

    }

    @Test
    public void TestDev4(){
        Solution sol = new Solution();


        int numCourses = 5;
        int[][] prerequisites = new int[][]{
                {1,4},{2,4},{3,1},{3,2}
        };

        boolean expect = true;

        boolean actual = sol.canFinish(numCourses, prerequisites);

        Assertions.assertEquals(expect, actual);

    }
}
