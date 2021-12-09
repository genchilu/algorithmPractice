import java.lang.reflect.Array;
import java.util.ArrayList;
import java.util.List;

class Solution {
    public List<List<Integer>> combine(int n, int k) {
        return combine(1,n,k);
    }

    public List<List<Integer>> combine(int begin, int n, int k) {
        var result = new ArrayList<List<Integer>>();
        if (k==1) {
            for(int i = begin;i<=n;i++){
                var tmp = new ArrayList<Integer>();
                tmp.add(i);
                result.add(tmp);
            }
        } else {
            for (int i = begin;i<=n;i++){
                var pres = combine(i+1,n,k-1);
                for(List<Integer> pre:pres) {
                    pre.add(i);
                    result.add(pre);
                }
            }
        }

        return result;
    }
}