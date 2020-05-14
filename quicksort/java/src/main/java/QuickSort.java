public class QuickSort {
    public static void sort(int[] input) {
        if(input!=null) {
            quickSort(input, 0, input.length - 1);
        }
    }

    private static void quickSort(int[] input, int begin, int end) {
        if (end>begin) {
            int pivotIdx = begin;
            for(int i = begin+1;i<=end;i++) {
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

    private static void swap(int[] input, int a, int b) {
        int tmp = input[a];
        input[a] = input[b];
        input[b] = tmp;
    }
}
