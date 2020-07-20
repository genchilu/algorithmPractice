import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class CountingInversions {

    public static int countInversions(int[] input) {

        if(input == null || input.length <=1) {
            return 0;
        }

        CountInversionsResult countInversionsResult = counting(input);

        return countInversionsResult.count;
    }

    private static CountInversionsResult counting(int[] input) {
        if(input.length <=1) {
            CountInversionsResult countInversionsResult = new CountInversionsResult();
            countInversionsResult.count = 0;
            countInversionsResult.sortedInput = input;

            return countInversionsResult;
        }

        int mid = input.length/2;

        CountInversionsResult lCountInversionsResult = counting(Arrays.copyOfRange(input, 0, mid));
        int lCount = lCountInversionsResult.count;
        int[] linput = lCountInversionsResult.sortedInput;

        CountInversionsResult rCountInversionsResult = counting(Arrays.copyOfRange(input, mid, input.length));
        int rCount = rCountInversionsResult.count;
        int[] rinput = rCountInversionsResult.sortedInput;

        int count = 0;
        int idx = 0;
        int lidx = 0;
        int ridx = 0;
        int[] sortedArray = new int[input.length];

        while (lidx < linput.length || ridx < rinput.length) {
            if (lidx == linput.length) {
                sortedArray[idx] = rinput[ridx];
                ridx++;
            } else if (ridx == rinput.length) {
                sortedArray[idx] = linput[lidx];
                lidx++;
            } else if (linput[lidx] <= rinput[ridx]) {
                sortedArray[idx] = linput[lidx];
                lidx++;
            } else {
                sortedArray[idx] = rinput[ridx];
                ridx++;
                count += linput.length-lidx;
            }
            idx++;
        }

        CountInversionsResult countInversionsResult = new CountInversionsResult();
        countInversionsResult.sortedInput = sortedArray;
        countInversionsResult.count = count+lCount+rCount;
        return countInversionsResult;
    }

}
