import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class CountingInversions {

    public static int CountInversions(int[] input) {
        if (input == null) {
            return 0;
        }

        if (input.length <=1) {
            return 0;
        }

        return CountWithMergeSort(input, 0, input.length).count;
    }

    private static CountInversionsResult CountWithMergeSort(int[] input, int begin, int end) {
        if((end-begin) == 1) {
            CountInversionsResult countInversionsResult = new CountInversionsResult();
            countInversionsResult.count = 0;
            countInversionsResult.sortedInput = Arrays.copyOfRange(input, begin, end);
            return countInversionsResult;
        } else {
            int mid = (end+begin)/2;
            CountInversionsResult lCountInversionsResult = CountWithMergeSort(input, begin, mid);
            CountInversionsResult rCountInversionsResult = CountWithMergeSort(input, mid, end);

            int i = 0;
            int j = 0;
            int count = 0;
            int[] sortedInput = new int[(end-begin)];
            int idx = 0;

            while (i < lCountInversionsResult.sortedInput.length || j < rCountInversionsResult.sortedInput.length) {
                if(i == lCountInversionsResult.sortedInput.length) {
                    sortedInput[idx] = rCountInversionsResult.sortedInput[j];
                    j++;
                } else if(j==rCountInversionsResult.sortedInput.length
                        || lCountInversionsResult.sortedInput[i] <= rCountInversionsResult.sortedInput[j]) {

                    sortedInput[idx] = lCountInversionsResult.sortedInput[i];
                    i++;
                } else {
                    sortedInput[idx] = rCountInversionsResult.sortedInput[j];
                    j++;
                    count += lCountInversionsResult.sortedInput.length - i;
                }
                idx ++;
            }

            CountInversionsResult countInversionsResult = new CountInversionsResult();
            countInversionsResult.sortedInput = sortedInput;
            countInversionsResult.count = count + lCountInversionsResult.count + rCountInversionsResult.count;

            return countInversionsResult;

        }
    }
}
