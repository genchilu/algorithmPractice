import java.util.HashMap;
import java.util.List;
import java.util.Map;

class Solution {
    public int minimumTotal(List<List<Integer>> triangle) {
        for(var i=triangle.size()-2;i>=0;i--){
            for(int j=0;j<triangle.get(i).size();j++){
                var val = triangle.get(i).get(j);
                var lm = triangle.get(i+1).get(j);
                var rm = triangle.get(i+1).get(j+1);
                triangle.get(i).set(j, val+Math.min(lm, rm));
            }
        }

        return triangle.get(0).get(0);
    }

}