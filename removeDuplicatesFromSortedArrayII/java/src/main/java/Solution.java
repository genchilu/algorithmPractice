class Solution {

    public int removeDuplicates(int[] nums) {
        int l = nums.length;
        int idx = 1;
        int pre = nums[0];
        int count = 1;

        while(idx<l) {
            if(nums[idx] == pre) {
                if (count==2) {
                    int cur = idx;
                    while(idx < l && nums[idx] == pre) {
                        idx++;
                    }
                    System.arraycopy(nums, idx, nums, cur, l-idx);
                    l = l - (idx-cur);
                    idx = cur;

                    count =1;
                    pre = nums[idx];
                } else {
                    count++;
                }
            } else {
                count =1;
                pre = nums[idx];
            }

            idx++;
        }

        return l;
    }
}