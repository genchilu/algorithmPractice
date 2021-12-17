import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    public void TestDev1(){

        //boolean isInterleave(String s1, String s2, String s3)
        Solution solution = new Solution();
        String s1="aabcc";
        String s2="dbbca";
        String s3="aadbbcbcac";

        boolean expect = true;

        var actual = solution.isInterleave(s1,s2,s3);

        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev2(){

        //boolean isInterleave(String s1, String s2, String s3)
        Solution solution = new Solution();
        String s1="aaaaa";
        String s2="bbbbb";
        String s3="aaaaabbbbb";

        boolean expect = true;

        var actual = solution.isInterleave(s1,s2,s3);

        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev3(){

        //boolean isInterleave(String s1, String s2, String s3)
        Solution solution = new Solution();
        String s1="aabcc";
        String s2="dbbca";
        String s3="aadbbbaccc";

        boolean expect = false;

        var actual = solution.isInterleave(s1,s2,s3);

        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev4(){

        //boolean isInterleave(String s1, String s2, String s3)
        Solution solution = new Solution();
        String s1="";
        String s2="";
        String s3="";

        boolean expect = true;

        var actual = solution.isInterleave(s1,s2,s3);

        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev5(){

        //boolean isInterleave(String s1, String s2, String s3)
        Solution solution = new Solution();
        String s1="bbbbbabbbbabaababaaaabbababbaaabbabbaaabaaaaababbbababbbbbabbbbababbabaabababbbaabababababbbaaababaa";
        String s2="babaaaabbababbbabbbbaabaabbaabbbbaabaaabaababaaaabaaabbaaabaaaabaabaabbbbbbbbbbbabaaabbababbabbabaab";
        String s3="babbbabbbaaabbababbbbababaabbabaabaaabbbbabbbaaabbbaaaaabbbbaabbaaabababbaaaaaabababbababaababbababbbababbbbaaaabaabbabbaaaaabbabbaaaabbbaabaaabaababaababbaaabbbbbabbbbaabbabaabbbbabaaabbababbabbabbab";

        boolean expect = false;

        var actual = solution.isInterleave(s1,s2,s3);

        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev6(){

        //boolean isInterleave(String s1, String s2, String s3)
        Solution solution = new Solution();
        String s1="aa";
        String s2="ab";
        String s3="abaa";

        boolean expect = true;

        var actual = solution.isInterleave(s1,s2,s3);

        Assertions.assertEquals(expect, actual);
    }
}
