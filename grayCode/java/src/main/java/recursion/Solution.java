package recursion;

import java.util.ArrayList;
import java.util.List;

class Solution {
    public List<Integer> grayCode(int n) {
        if (n==0) {
            return new ArrayList<>();
        }

        if (n==1) {
            var result = new ArrayList<Integer>();
            result.add(0);
            result.add(1);
            return result;
        }


        var result = grayCode(n-1);

        for (int i = result.size()-1;i>=0;i--){
            var ele = result.get(i);
            ele += (1<<(n-1));
            result.add(ele);
        }

        return result;
    }
}



