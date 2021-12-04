package recursion;

import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

import java.util.ArrayList;

public class SolutionTest {

    @Test
    public void TestDev1(){
        var solution = new Solution();
        var input = 0;

        var expect = new ArrayList<Integer>();

        var actual = solution.grayCode(input);

        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev2(){
        var solution = new Solution();
        var input = 1;

        var expect = new ArrayList<Integer>();
        expect.add(0);
        expect.add(1);

        var actual = solution.grayCode(input);

        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev3(){
        var solution = new Solution();
        var input = 2;

        var expect = new ArrayList<Integer>();
        expect.add(0);
        expect.add(1);
        expect.add(3);
        expect.add(2);

        var actual = solution.grayCode(input);

        Assertions.assertEquals(expect, actual);
    }

    @Test
    public void TestDev4(){
        var solution = new Solution();
        var input = 3;

        var expect = new ArrayList<Integer>();
        expect.add(0);
        expect.add(1);
        expect.add(3);
        expect.add(2);
        expect.add(6);
        expect.add(7);
        expect.add(5);
        expect.add(4);

        var actual = solution.grayCode(input);

        Assertions.assertEquals(expect, actual);
    }
}
