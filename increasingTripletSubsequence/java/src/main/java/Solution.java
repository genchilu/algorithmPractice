import java.util.ArrayList;
import java.util.Collections;

class Solution {
    public boolean increasingTriplet(int[] nums) {
        ArrayList<Integer> inc = new ArrayList<>();
        for(int num:nums) {
            var idx = Collections.binarySearch(inc, num);
            if (idx<0) {
                idx = idx+1;
                idx = -idx;
                if (idx==inc.size()) {
                    inc.add(num);
                } else {
                    inc.set(idx, num);
                }
            }
        }
        return inc.size()>=3;
    }
}