public class QuickSort {
    public static void sort(int[] input) {
        if (input!=null) {
            quicksort(input, 0, input.length);
        }
    }

    private static void quicksort(int[] input, int begin, int end) {
        if((end-begin) > 1) {
            int pivotIdx = begin;

            swap(input, pivotIdx, begin);
            pivotIdx = begin;

            for(int i=begin+1;i<end;i++){
                if(input[i] < input[begin]){
                    pivotIdx++;
                    swap(input, i, pivotIdx);
                }
            }

            swap(input, begin, pivotIdx);
            quicksort(input, begin, pivotIdx);
            quicksort(input, pivotIdx+1, end);
        }
    }

    private static void swap(int[] input, int from, int to) {
        if(from != to) {
            int tmp = input[from];
            input[from] = input[to];
            input[to] = tmp;
        }
    }

}
