import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    public void TestDev1(){
        String start = "AACCGGTT";
        String end = "AACCGGTA";

        String[] bank = new String[]{"AACCGGTA"};

        int expect = 1;

        var sol = new Solution();

        var actual = sol.minMutation(start, end,bank);

        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev2(){
        String start = "AACCGGTT";
        String end = "AAACGGTA";

        String[] bank = new String[]{"AACCGGTA","AACCGCTA","AAACGGTA"};

        int expect = 2;

        var sol = new Solution();

        var actual = sol.minMutation(start, end,bank);

        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev3(){
        String start = "AAAAACCC";
        String end = "AACCCCCC";

        String[] bank = new String[]{"AAAACCCC","AAACCCCC","AACCCCCC"};

        int expect = 3;

        var sol = new Solution();

        var actual = sol.minMutation(start, end,bank);

        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev4(){
        String start = "AAAAACCC";
        String end = "AACCCCCA";

        String[] bank = new String[]{"AAAACCCC","AAACCCCC","AACCCCCC"};

        int expect = -1;

        var sol = new Solution();

        var actual = sol.minMutation(start, end,bank);

        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev5(){
        String start = "AACCGGTT";
        String end = "AAACGGTA";

        String[] bank = new String[]{"AACCGGTA","AACCGCTA","AAACGGTA"};

        int expect = 2;

        var sol = new Solution();

        var actual = sol.minMutation(start, end,bank);

        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev6(){
        String start = "AAAACCCC";
        String end = "CCCCCCCC";

        String[] bank = new String[]{"AAAACCCA","AAACCCCA","AACCCCCA","AACCCCCC","ACCCCCCC","CCCCCCCC","AAACCCCC","AACCCCCC"};

        int expect = 4;

        var sol = new Solution();

        var actual = sol.minMutation(start, end,bank);

        Assertions.assertEquals(expect, actual);
    }
}
