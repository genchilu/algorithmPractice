import java.util.ArrayList;
import java.util.List;

class Solution {
    public List<List<Integer>> findRLEArray(int[][] encoded1, int[][] encoded2) {
        var result = new ArrayList<List<Integer>>();

        int i = 0;
        int j = 0;

        while(i<encoded1.length) {
            while (j<encoded2.length){
                var e1 = encoded1[i];
                var v1 = e1[0];
                var c1 = e1[1];
                var e2 = encoded2[j];
                var v2 = e2[0];
                var c2 = e2[1];

                var min = Math.min(c1,c2);

                var newv = v1*v2;
                if(result.size()>0 &&  result.get(result.size()-1).get(0).equals(newv)) {
                    var encode = result.get(result.size()-1);
                    var count = encode.get(1) + min;
                    encode.remove(1);
                    encode.add(count);

                    result.remove(result.size()-1);
                    result.add(encode);
                } else {
                    var encode = new ArrayList<Integer>();
                    encode.add(newv);
                    encode.add(min);
                    result.add(encode);
                }

                e1[1]-=min;
                if(e1[1] == 0 ) {
                    i++;
                }
                e2[1]-=min;
                if(e2[1] == 0 ) {
                    j++;
                }
            }
        }

        return result;
    }
}
