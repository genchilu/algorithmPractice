def sort(input):
    if input is not None:
        _quick_sort(input, 0, len(input))

def _quick_sort(input, begin, end):
    if (end-begin) > 1:
        pivot_idx = begin

        _swap(input, pivot_idx, begin)

        final_pivot_idx = begin
        for i in range(begin+1, end):
            if input[i] < input[pivot_idx]:
                final_pivot_idx = final_pivot_idx+1
                _swap(input, i, final_pivot_idx)

        _swap(input, begin, final_pivot_idx)

        _quick_sort(input, begin, final_pivot_idx)
        _quick_sort(input, final_pivot_idx+1, end)

def _swap(input, from_idx, to_idx):
    if(from_idx != to_idx):
        tmp = input[from_idx]
        input[from_idx]= input[to_idx]
        input[to_idx] = tmp
        