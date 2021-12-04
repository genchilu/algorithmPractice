package iteration;

import java.util.ArrayList;
import java.util.List;

class Solution {
    public List<Integer> grayCode(int n) {
        if (n==0) {
            return new ArrayList<>();
        }


        var result = new ArrayList<Integer>();
        result.add(0);
        result.add(1);

        for (int i=2;i<=n;i++){
            for (int j=result.size()-1;j>=0;j--){
                var ele = result.get(j);
                ele+=(1<<(i-1));
                result.add(ele);
            }
        }

        return result;
    }
}



