def countInversions(input):
    if input == None:
        return 0
    count, sort_inpiut = _count(input)

    return count

def _count(input):
    if (len(input)<2):
        return 0, input
    mid = len(input)//2
    lcount, linput = _count(input[0:mid])
    rcount, rinput = _count(input[mid:])

    lidx = 0
    ridx = 0
    count = 0
    sorted_input = []

    while lidx < len(linput) or ridx < len(rinput):
        if (lidx == len(linput)):
            sorted_input.extend(rinput[ridx:])
            ridx = len(rinput)
        elif (ridx == len(rinput)):
            sorted_input.extend(linput[lidx:])
            lidx = len(linput)
        elif (linput[lidx] <= rinput[ridx]):
            sorted_input.append(linput[lidx])
            lidx += 1
        else:
            sorted_input.append(rinput[ridx])
            ridx += 1
            count += len(linput) - lidx
        
    return (count+lcount+rcount), sorted_input

