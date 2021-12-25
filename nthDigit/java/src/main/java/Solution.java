class Solution {
    public int findNthDigit(int n) {

        int digitNum = 1;
        long range = 9;
        while(n>(range*digitNum)) {
            n-=range*digitNum;
            range*=10;
            digitNum++;
        }

        int startNo = (int)(range/9);
        int idx = (n-1)/digitNum;   // 1/2=0
        int reminder = (n-1)%digitNum; //1

        int res = startNo+idx;  //10
        int div = 1;

        for(int i=1;i<digitNum;i++){
            div*=10;
        }
        //div=10

        for(int i =0;i<reminder;i++){
            res%=div;
            div/=10;
        }

        return res/div;
//        1~9 9*1
//        10~99 90*2
//        100~999 900*3
    }
}

//189 99
//190 100
//191 100
//192 100
//193 101
//194 101
//195 101
//196 102
//197 102
//198 102