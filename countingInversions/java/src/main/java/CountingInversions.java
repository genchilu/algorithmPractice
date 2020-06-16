import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class CountingInversions {

    public static int CountInversions(int[] input) {
        if (input == null) {
            return 0;
        }

        return counting(input).count;
    }

    private static CountInversionsResult counting(int[] input) {
        if(input.length <= 1) {
            CountInversionsResult countInversionsResult = new CountInversionsResult();
            countInversionsResult.count = 0;
            countInversionsResult.sortedInput = input;
            return countInversionsResult;
        }

        int mid = input.length/2;

        CountInversionsResult lCountInversionsResult =counting(Arrays.copyOfRange(input, 0, mid));
        CountInversionsResult rCountInversionsResult =counting(Arrays.copyOfRange(input, mid, input.length));

        int count =0;
        int idx = 0;
        int lidx = 0;
        int ridx = 0;
        int[] mergedSortedInput = new int[input.length];
        while (lidx < lCountInversionsResult.sortedInput.length || ridx < rCountInversionsResult.sortedInput.length) {
            if(lidx == lCountInversionsResult.sortedInput.length) {
                mergedSortedInput[idx] = rCountInversionsResult.sortedInput[ridx];
                idx++;
                ridx++;
            } else if (ridx == rCountInversionsResult.sortedInput.length) {
                mergedSortedInput[idx] = lCountInversionsResult.sortedInput[lidx];
                idx++;
                lidx++;
            } else if (lCountInversionsResult.sortedInput[lidx] <= rCountInversionsResult.sortedInput[ridx]) {
                mergedSortedInput[idx] = lCountInversionsResult.sortedInput[lidx];
                idx++;
                lidx++;
            } else {
                mergedSortedInput[idx] = rCountInversionsResult.sortedInput[ridx];
                idx++;
                ridx++;
                count += lCountInversionsResult.sortedInput.length - lidx;
            }
        }

        CountInversionsResult countInversionsResult = new CountInversionsResult();
        countInversionsResult.sortedInput = mergedSortedInput;
        countInversionsResult.count = count+lCountInversionsResult.count+rCountInversionsResult.count;

        return countInversionsResult;
    }
}
