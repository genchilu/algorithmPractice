class Solution {
    public int nthSuperUglyNumber(int n, int[] primes) {
        int[] res = new int[n];
        res[0]=1;

        int[] candidatesIdx = new int[primes.length];
        for(int i = 1;i<n;i++) {
            int min = Integer.MAX_VALUE;
            for (int j = 0; j < candidatesIdx.length; j++) {
                int idx = candidatesIdx[j];
                int candidate = res[idx] * primes[j];
                if (candidate < min) {
                    min = candidate;
                }
            }

            res[i] = min;

            for (int j = 0; j < candidatesIdx.length; j++) {
                int idx = candidatesIdx[j];
                int candidate = res[idx] * primes[j];
                if (candidate == min) {
                    candidatesIdx[j]++;
                }
            }
        }



        return res[n-1];
    }
}