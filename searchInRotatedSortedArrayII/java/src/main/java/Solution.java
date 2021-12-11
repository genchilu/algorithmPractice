class Solution {
    public boolean search(int[] nums, int target) {
        var pivot = findPivot(nums);
        var l =0;
        var r = nums.length-1;
        while(l<=r) {
            int m = (l+r)/2;

            var mapm =(m+pivot)%nums.length;

            if(nums[mapm]==target) {
                return true;
            }

            if(nums[mapm] > target) {
                r=m-1;
            } else {
                l=m+1;
            }
        }
        return false;
    }

    public int findPivot(int[] nums) {

        if(nums.length==1) {
            return 0;
        }

        if(nums[0] == nums[nums.length-1]) {
            for(int i = 1;i<nums.length;i++){
                if (nums[i] < nums[i-1]) {
                    return i;
                }
            }

            return 0;
        }

        int l = 0;
        int r = nums.length-1;
        while (l<=r) {
            int m = (l+r)/2;
            if (m == nums.length - 1) {
                break;
            }
            if(nums[m] > nums[m+1]) {
                return m+1;
            }
            if (nums[m]>= nums[0]) {
                l = m+1;
            } else if (nums[m] < nums[0]){
                r = m-1;
            }
        }

        return (l+1)%nums.length;
    }
}