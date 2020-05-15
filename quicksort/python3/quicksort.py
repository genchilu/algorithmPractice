def sort(input):
    if input is not None:
        _quicksort(input, 0, len(input))

def _quicksort(input, begin, end):
    if end > begin:
        pivot = begin
        for i in range(begin, end):
            if input[i]<input[begin]:
                pivot=pivot+1
                _swap(input, i, pivot)
        
        _swap(input, begin, pivot)
        _quicksort(input, begin, pivot)
        _quicksort(input, pivot+1, end)

def _swap(input, a, b):
    tmp=input[a]
    input[a]=input[b]
    input[b]=tmp
        