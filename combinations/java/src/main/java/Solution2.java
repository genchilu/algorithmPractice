import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;

public class Solution2 {
    public List<List<Integer>> combine(int n, int k) {

        if (k==0 || n==0 || k>n) {
            return Collections.emptyList();
        }
        List<List<Integer>> combs = new ArrayList<>();
        for(int i =1;i<=n;i++){
            combs.add(Arrays.asList(i));
        }

        for(int i=2;i<=k;i++){
            List<List<Integer>> newCombs = new ArrayList<>();
            for(int j=i;j<=n;j++){
                for(var comb:combs) {
                    if(comb.get(comb.size()-1) < j) {
                        var newComb = new ArrayList<>(comb);
                        newComb.add(j);
                        newCombs.add(newComb);
                    }
                }
            }
            combs = newCombs;
        }
        return combs;
    }


}
