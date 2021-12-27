import java.util.ArrayList;
import java.util.List;

class Solution {
    public List<Integer> findDuplicates(int[] nums) {
        List<Integer> res = new ArrayList<>();

        for(int num:nums){
            if(num<0) {
                num = -num;
            }

            if(nums[num-1]<0) {
                res.add(num);
            } else {
                nums[num-1] = -nums[num-1];
            }
        }

        return res;
    }
}
