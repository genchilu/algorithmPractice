import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

import java.util.ArrayList;
import java.util.List;

public class SolutionTest {

    @Test
    public void TestDev1(){
        //findRLEArray(int[][] encoded1, int[][] encoded2) {
        var solution = new Solution();
        var input1 = new  int[][]{{1,3},{2,3}};
        var input2 = new  int[][]{{6,3},{3,3}};

        var expect = new ArrayList<List<Integer>>();
        var encode1 = new ArrayList<Integer>();
        encode1.add(6);
        encode1.add(6);
        expect.add(encode1);

        var actual =solution.findRLEArray(input1, input2);

        Assertions.assertEquals(expect,actual);
    }

    @Test
    public void TestDev2(){
        //findRLEArray(int[][] encoded1, int[][] encoded2) {
        var solution = new Solution();
        var input1 = new  int[][]{{1,3},{2,1},{3,2}};
        var input2 = new  int[][]{{2,3},{3,3}};

        var expect = new ArrayList<List<Integer>>();
        var encode1 = new ArrayList<Integer>();
        encode1.add(2);
        encode1.add(3);
        var encode2 = new ArrayList<Integer>();
        encode2.add(6);
        encode2.add(1);
        var encode3 = new ArrayList<Integer>();
        encode3.add(9);
        encode3.add(2);

        expect.add(encode1);
        expect.add(encode2);
        expect.add(encode3);

        var actual =solution.findRLEArray(input1, input2);

        Assertions.assertEquals(expect,actual);
    }

    @Test
    public void TestDev3(){
        //findRLEArray(int[][] encoded1, int[][] encoded2) {
        var solution = new Solution();
        var input1 = new  int[][]{{1,20},{19,14},{1,58}};
        var input2 = new  int[][]{{15,42},{5,50}};

        var expect = new ArrayList<List<Integer>>();
        var encode1 = new ArrayList<Integer>();
        encode1.add(15);
        encode1.add(20);
        var encode2 = new ArrayList<Integer>();
        encode2.add(285);
        encode2.add(14);
        var encode3 = new ArrayList<Integer>();
        encode3.add(15);
        encode3.add(8);
        var encode4 = new ArrayList<Integer>();
        encode4.add(5);
        encode4.add(50);

        expect.add(encode1);
        expect.add(encode2);
        expect.add(encode3);
        expect.add(encode4);

        var actual =solution.findRLEArray(input1, input2);

        Assertions.assertEquals(expect,actual);
    }
}
