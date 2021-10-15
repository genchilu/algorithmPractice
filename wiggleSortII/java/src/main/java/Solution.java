class Solution {
    public void wiggleSort(int[] nums) {
        var midVal = findKth(nums,0,nums.length-1, (0+nums.length)/2);

        var l = 0;
        var r = nums.length-1;
        var i = 0;
        while (i<=r) {
            var newIdx = getNewIdx(nums, i);

            if (nums[newIdx] < midVal) {
                swap(nums, newIdx, getNewIdx(nums,r));
                r--;
            } else if (nums[newIdx] > midVal) {
                swap(nums, newIdx, getNewIdx(nums,l));
                l++;
                i++;
            } else {
                i++;
            }
        }
    }

    public static int getNewIdx(int[] nums, int i) {
        return  (1+2*i) % (nums.length | 1);
    }

    public static int findKth(int[] nums, int l, int r, int k) {
        if(nums.length==1) {
            return nums[0];
        }

        var idx = l+1;
        for(int i=l+1;i<=r;i++) {
            if (nums[i]>nums[l]) {
                swap(nums, i, idx);
                idx++;
            }
        }

        swap(nums, l, idx-1);
        if (idx==k) {
            return nums[idx-1];
        } else if (idx > k) {
            return findKth(nums, l, idx-2,k);
        } else {
            return findKth(nums, idx, r, k);
        }

    }

    public static void swap(int[] nums, int a, int b) {
        var tmp = nums[a];
        nums[a] = nums[b];
        nums[b] = tmp;
    }
}