public class QuickSort {
    public static void sort(int[] input) {
        if (input != null) {
            quicksort(input, 0, input.length);
        }

    }

    private static int[] quicksort(int[] input, int begin, int end) {
        if ((end-begin) > 1) {
            int pivotIdx = begin;

            swap(input, begin, pivotIdx);
            int finalPivotIdx = begin;

            for (int i = begin+1; i<end;i++){
                if(input[i] < input[pivotIdx]) {
                    swap(input, i, finalPivotIdx+1);
                    finalPivotIdx++;
                }
            }

            swap(input, begin, finalPivotIdx);

            quicksort(input, begin, finalPivotIdx);
            quicksort(input, finalPivotIdx+1, end);
        }

        return input;
    }

    private static  void swap(int[] input, int fromIdx, int toIdx) {
        if(fromIdx!=toIdx) {
            int tmp = input[fromIdx];
            input[fromIdx] = input[toIdx];
            input[toIdx] = tmp;
        }
    }
}
