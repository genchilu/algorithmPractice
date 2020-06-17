def countInversions(input):
    if input == None:
        return 0
    [count, sorted_array] = counting(input)
    return count

def counting(input):
    if len(input) <= 1:
        return 0, input
    
    mid = len(input)//2
    lcount, linput = counting(input[0: mid])
    rcount, rinput = counting(input[mid: len(input)])

    count = 0
    lidx = 0
    ridx = 0
    merged_input = []

    while lidx<len(linput) or ridx <len(rinput):
        if lidx == len(linput):
            merged_input.append(rinput[ridx])
            ridx = ridx +1
        elif ridx == len(rinput):
            merged_input.append(linput[lidx])
            lidx = lidx +1
        elif linput[lidx] <= rinput[ridx]:
            merged_input.append(linput[lidx])
            lidx = lidx +1
        else:
            merged_input.append(rinput[ridx])
            ridx = ridx +1
            count = count + len(linput) - lidx
        
    return (count+lcount+rcount), merged_input

