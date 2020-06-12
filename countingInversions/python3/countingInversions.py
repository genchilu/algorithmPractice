def countInversions(input):
    if input == None:
        return 0
    if len(input) <=1:
        return 0
    [sorted_array, count] = countWithMergeSort(input)
    return count

def countWithMergeSort(input):
    if len(input) == 1:
        return input, 0

    mid = len(input)//2
    [lsorted_array, lcount] = countWithMergeSort(input[0: mid])
    [rsorted_array, rcount] = countWithMergeSort(input[mid: len(input)])

    i=0
    j=0
    count = 0
    sorted_array = []

    while i != len(lsorted_array) or j != len(rsorted_array):
        if i == len(lsorted_array):
            sorted_array.append(rsorted_array[j])
            j = j + 1
        elif j == len(rsorted_array) or lsorted_array[i] <= rsorted_array[j]:
            sorted_array.append(lsorted_array[i])
            i = i + 1
        else:
            sorted_array.append(rsorted_array[j])
            j = j + 1
            count = count + len(lsorted_array) - i

    return sorted_array, (count + lcount + rcount)

