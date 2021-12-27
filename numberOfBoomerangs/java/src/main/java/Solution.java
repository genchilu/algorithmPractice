class Solution {
    public int numberOfBoomerangs(int[][] points) {

        int res = 0;
        for(int i =0;i<points.length;i++){
            int [] a= points[i];
            for(int j=0;j<points.length;j++) {
                if(j!=i) {
                    int [] b= points[j];
                    int ab = (a[0]-b[0]) * (a[0]-b[0]) + (a[1]-b[1]) * (a[1]-b[1]);
                    for (int k = j + 1; k < points.length; k++) {
                        if(k!=i){
                            int [] c= points[k];
                            int ac = (a[0]-c[0]) * (a[0]-c[0]) + (a[1]-c[1]) * (a[1]-c[1]);
                            if(ab==ac) {
                                res+=2;
                            }
                        }
                    }
                }
            }
        }
        return res;
    }
}

//