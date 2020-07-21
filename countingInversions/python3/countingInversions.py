def countInversions(input):
    if input == None or len(input) <=1:
        return 0

    [sorted_array, count] = counting(input)

    return count


def counting(input):
    if len(input) <= 1:
        return input, 0
    
    mid = len(input)//2
    [linput, lcount] = counting(input[0:mid])
    [rinput, rcount] = counting(input[mid:])

    count = 0
    lidx = 0
    ridx =0
    sort_array = []

    while lidx < len(linput) or ridx < len(rinput):
        if lidx == len(linput):
            sort_array.extend(rinput[ridx:])
            ridx = len(rinput)
        elif ridx == len(rinput):
            sort_array.extend(linput[lidx:])
            lidx = len(linput)
        elif linput[lidx] <= rinput[ridx]:
            sort_array.append(linput[lidx])
            lidx = lidx+1
        else:
            sort_array.append(rinput[ridx])
            ridx = ridx+1
            count = count + len(linput) - lidx
    
    return sort_array, count+lcount+rcount
