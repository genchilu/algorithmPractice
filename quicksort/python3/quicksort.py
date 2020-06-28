def sort(input):
    if input is not None:
        _quicksort(input, 0, len(input))


def _quicksort(input, begin, end):
    if (end-begin) > 1:
        pivotIdx = begin
        
        _swap(input, pivotIdx, begin)
        pivotIdx = begin

        for i in range(begin+1, end):
            if input[i] < input[begin]:
                pivotIdx = pivotIdx +1
                _swap(input, pivotIdx, i)
        
        _swap(input, begin, pivotIdx)
        _quicksort(input, begin, pivotIdx)
        _quicksort(input, pivotIdx+1, end)


def _swap(input, a, b):
    if a != b:
        tmp = input[a]
        input[a] = input[b]
        input[b] = tmp

        