def sort(input):
    if input != None and len(input) > 1:
        _quicksort(input, 0, len(input))

def _quicksort(input, begin, end):
    if (end-begin) > 1:
        pivot = begin
        _swap(input, pivot, begin)
        pivot = begin

        for i in range(begin+1, end):
            if input[i] < input[begin]:
                pivot = pivot+1
                _swap(input, pivot, i)
        
        _swap(input, begin, pivot)

        _quicksort(input, begin, pivot)
        _quicksort(input, pivot+1, end)

def _swap(input, from_idx, to_idx):
    if from_idx != to_idx:
        tmp = input[from_idx]
        input[from_idx] = input[to_idx]
        input[to_idx] = tmp

        