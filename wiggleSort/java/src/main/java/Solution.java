class Solution {
    public void wiggleSort(int[] nums) {
        if(nums.length<=1) {
            return;
        }

        for(int i =1;i<nums.length;i+=2){
            if(i==nums.length-1) {
                if(nums[i]<nums[i-1]) {
                    swap(nums, i, i-1);
                }
            }else {
                if(nums[i-1] >= nums[i] && nums[i-1] >= nums[i+1]) {
                    swap(nums, i-1,i);
                } else if (nums[i+1]>=nums[i] && nums[i+1] >= nums[i-1]) {
                    swap(nums, i+1, i);
                }
            }
        }
    }

    public void swap(int[] nums, int a, int b){
        var tmp = nums[a];
        nums[a] = nums[b];
        nums[b] = tmp;
    }
}