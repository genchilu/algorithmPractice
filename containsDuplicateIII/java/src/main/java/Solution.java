import java.math.BigInteger;
import java.util.HashMap;
import java.util.Map;

class Solution {
    public boolean containsNearbyAlmostDuplicate(int[] nums, int k, int t) {
        Map<Long, Long> bucket = new HashMap<>();
        for(int i=0;i<nums.length;i++){

            long label = getLabel(nums[i],t);
            if(bucket.containsKey(label)) {
                return true;
            }
            if(bucket.containsKey(label-1) && nums[i]-bucket.get(label-1) <= t) {
                return true;
            }
            if(bucket.containsKey(label+1) && bucket.get(label+1)-nums[i] <= t) {
                return true;
            }

            bucket.put(label, (long)nums[i]);

            if(bucket.size()>k) {
                long removeLabel = getLabel(nums[i-k],t);
                bucket.remove(removeLabel);
            }

        }


        return false;
    }

    public long getLabel(int val, int t) {
        long remap = ((long)val- (long)Integer.MIN_VALUE);
        long label = remap/((long)t+1);
        return label;

    }
}