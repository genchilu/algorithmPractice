def do_kosaraju_scc(edges):
    g = {}
    for edge in edges:
        if edge[0] not in g:
            g[edge[0]] = []
        g[edge[0]].append(edge[1])

        if edge[1] not in g:
            g[edge[1]] = []
    
    is_vistited = {}
    for v in g.keys():
        is_vistited[v] = False
    
    finish_stack = []
    for v in g.keys():
        dfs_round1(g, v, is_vistited, finish_stack)

    for v in g.keys():
        is_vistited[v] = False

    rev_g = reverse_g(g)
    componments = {}
    while(len(finish_stack) > 0):
        v = finish_stack.pop()
        componment = []
        dfs_round2(rev_g, v, is_vistited, componment)

        if(len(componment) > 0):
            componment.sort()
            componments[componment[0]] = componment

    return componments

def dfs_round1(g, v, is_vistited, finish_stack):
    if not is_vistited[v]:
        is_vistited[v] = True
        for to_v in g[v]:
            dfs_round1(g, to_v, is_vistited, finish_stack)
        finish_stack.append(v)

def reverse_g(g):
    rev_g = {}
    for from_v in g.keys():
        for to_v in g[from_v]:
            if to_v not in rev_g:
                rev_g[to_v] = []
            rev_g[to_v].append(from_v)

        if from_v not in rev_g:
            rev_g[from_v] = []
    
    return rev_g

def dfs_round2(g, v, is_vistited, componment):
    if not is_vistited[v]:
        is_vistited[v] = True
        for to_v in g[v]:
            dfs_round2(g, to_v, is_vistited, componment)
        componment.append(v)

