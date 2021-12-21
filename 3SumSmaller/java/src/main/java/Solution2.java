import java.util.Arrays;

//-2 0 1 3
//2
//-2 0 3
//-2 0 1

class Solution2 {
    public int threeSumSmaller(int[] nums, int target) {
        Arrays.sort(nums);
        int i =0;
        int j = i+1;
        int k = nums.length-1;

        int res = 0;
        while(i<k){
            while(j<k) {
                int sum = nums[i] + nums[j]+nums[k];
                if(sum<target) {
                    res+=k-j;
                    j++;
                } else {
                    k--;
                }
            }
            i++;
            j=i+1;
            k = nums.length-1;
        }
        return res;
    }
}