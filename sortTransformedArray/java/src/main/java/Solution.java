class Solution {
    public int[] sortTransformedArray(int[] nums, int a, int b, int c) {

        int[] res = new int[nums.length];
        if(a>0){
            int i=0;
            int j =nums.length-1;
            int idx=res.length-1;
            while(i<=j){
                int l = calculate(a,b,c,nums[i]);
                int r = calculate(a,b,c,nums[j]);
                if(l<r){
                    res[idx]=r;
                    j--;
                } else {
                    res[idx]=l;
                    i++;
                }
                idx--;
            }
        } else {
            int i=0;
            int j =nums.length-1;
            int idx=0;
            while(i<=j){
                int l = calculate(a,b,c,nums[i]);
                int r = calculate(a,b,c,nums[j]);
                if(l<r){
                    res[idx]=l;
                    i++;
                } else {
                    res[idx]=r;
                    j--;
                }
                idx++;
            }
        }
        return res;
    }

    public int calculate(int a, int b, int c, int x) {
        return a*x*x+b*x+c;
    }
}

