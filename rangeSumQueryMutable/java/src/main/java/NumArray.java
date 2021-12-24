class NumArray {

    int[] nums;
    int[] chunks;
    int chunkSize;
    int chunkLen;
    public NumArray(int[] nums) {
        this.nums = nums;

        double l = Math.sqrt((double)(nums.length));
        this.chunkSize = (int)Math.ceil(l);

        this.chunkLen = (nums.length/chunkSize)+1;
        this.chunks = new int[chunkLen];

        for(int i =0;i<this.chunkLen;i++){
            for(int j = 0;j<this.chunkSize;j++){
                if ((i*this.chunkSize+j)<nums.length) {
                    this.chunks[i]+=nums[i*this.chunkSize+j];
                }
            }
        }
    }


    public void update(int index, int val) {

        int chunkIdx = index/this.chunkSize;
        this.chunks[chunkIdx]-=this.nums[index];
        this.chunks[chunkIdx]+=val;

        nums[index] = val;
    }

    public int sumRange(int left, int right) {

        int lchunkIdx = left/this.chunkSize;
        int rchunkIdx = right/this.chunkSize;
        int sum = 0;
        for(int i = lchunkIdx;i<=rchunkIdx;i++){
            sum+=this.chunks[i];
        }

        for(int i = lchunkIdx*this.chunkSize;i<left;i++){
            sum-=this.nums[i];
        }

        for(int i = right+1;i<(rchunkIdx+1)*this.chunkSize;i++){
            if (i<this.nums.length) {
                sum -= this.nums[i];
            }
        }

        return sum;
    }
}

/**
 * Your NumArray object will be instantiated and called as such:
 * NumArray obj = new NumArray(nums);
 * obj.update(index,val);
 * int param_2 = obj.sumRange(left,right);
 */