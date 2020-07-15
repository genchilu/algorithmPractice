public class QuickSort {
    public static void sort(int[] input) {
        if(input != null && input.length > 1) {
            quickSort(input, 0, input.length);
        }
    }

    private static void quickSort(int[] input, int begin, int end) {
        if ((end-begin) > 1) {
            int pivotIdx = begin;

            swap(input, pivotIdx, begin);

            pivotIdx = begin;

            for (int i = (begin+1);i<end;i++) {
                if(input[i] < input[begin]) {
                    pivotIdx++;
                    swap(input, pivotIdx, i);
                }
            }

            swap(input, pivotIdx, begin);

            quickSort(input, begin, pivotIdx);
            quickSort(input, pivotIdx+1, end);
        }
    }
    private static void swap(int[] input, int from, int to) {
        if (from!=to) {
            int tmp = input[from];
            input[from] = input[to];
            input[to] = tmp;
        }
    }
}
