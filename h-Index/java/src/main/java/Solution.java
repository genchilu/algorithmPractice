import java.util.Arrays;
import java.util.Comparator;

class Solution {
    public int hIndex(int[] citations) {
        Arrays.sort(citations);


        for(int i = citations.length-1;i>=0;i++){
            int count = citations.length-i;
            if(count==citations[i]) {
                return count;
            } else if(count > citations[i]) {
                return count-1;
            }
        }

        return citations.length;
    }
}
