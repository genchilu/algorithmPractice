import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class CountingInversions {

    public static int CountInversions(int[] input) {
        if (input == null) {
            return 0;
        }

        CountInversionsResult countInversionsResult = count(input);
        return countInversionsResult.count;
    }

    private static CountInversionsResult count(int[] input) {
        if(input.length <=1) {
            CountInversionsResult countInversionsResult = new CountInversionsResult();
            countInversionsResult.count = 0;
            countInversionsResult.sortedInput = input;
            return countInversionsResult;
        }

        int mid = input.length/2;
        CountInversionsResult lCountInversionsResult = count(Arrays.copyOfRange(input, 0, mid));
        CountInversionsResult rCountInversionsResult = count(Arrays.copyOfRange(input, mid, input.length));

        int lidx = 0;
        int ridx = 0;
        int count = 0;
        int idx =0;
        int[] sortedArray = new int[input.length];

        while (lidx<lCountInversionsResult.sortedInput.length || ridx<rCountInversionsResult.sortedInput.length) {
            if(lidx == lCountInversionsResult.sortedInput.length) {
                sortedArray[idx] = rCountInversionsResult.sortedInput[ridx];
                ridx++;
            } else if (ridx == rCountInversionsResult.sortedInput.length) {
                sortedArray[idx] = lCountInversionsResult.sortedInput[lidx];
                lidx++;
            } else if (lCountInversionsResult.sortedInput[lidx] <= rCountInversionsResult.sortedInput[ridx]) {
                sortedArray[idx] = lCountInversionsResult.sortedInput[lidx];
                lidx++;
            } else {
                sortedArray[idx] = rCountInversionsResult.sortedInput[ridx];
                ridx++;
                count+=(lCountInversionsResult.sortedInput.length - lidx);
            }
            idx++;
        }

        CountInversionsResult countInversionsResult = new CountInversionsResult();
        countInversionsResult.sortedInput = sortedArray;
        countInversionsResult.count = lCountInversionsResult.count + rCountInversionsResult.count + count;

        return countInversionsResult;


    }

}
